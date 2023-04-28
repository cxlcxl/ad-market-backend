package servicepayment

import (
	"crypto/rand"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	utils2 "github.com/wechatpay-apiv3/wechatpay-go/utils"
	"market/app/model"
	"market/app/validator/v_data"
	"market/app/vars"
)

const (
	// NonceSymbols 随机字符串可用字符集
	NonceSymbols = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// NonceLength 随机字符串的长度
	NonceLength = 32
)

type ActionResponse struct {
	OutTradeNo    string `json:"out_trade_no"`
	Attach        string `json:"attach"`
	TransactionId string `json:"transaction_id"`
	TradeState    string `json:"trade_state"` // SUCCESS 表示成功
	BankType      string `json:"bank_type"`
	Payer         struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	SceneInfo struct {
		DeviceId string `json:"device_id"`
	} `json:"scene_info"`
}

// UserOrder 用户下单
func UserOrder(ctx *gin.Context, mobile string) (h5Url, outTradeNo string, err error) {
	order, err := buildLocalOrder(mobile)
	if err != nil {
		return "", "", err
	}
	// 微信提供的包
	outTradeNo = order.OutTradeNo
	h5Url, err = wxPayOrder(ctx, order.OutTradeNo, ctx.ClientIP(), order.OpenId, int64(order.Amt))
	fmt.Println(err)
	return
}

// JsApiOrder 用户下单
func JsApiOrder(ctx *gin.Context, mobile string) (h5Url, outTradeNo string, err error) {
	order, err := buildLocalOrder(mobile)
	if err != nil {
		return "", "", err
	}
	// 微信提供的包
	outTradeNo = order.OutTradeNo
	h5Url, err = wxPayJsApiOrder(ctx, order.OutTradeNo, ctx.ClientIP(), order.OpenId, int64(order.Amt))
	fmt.Println(err)
	return
}

// OrderQuery 用户订单查询
func OrderQuery(ctx *gin.Context, outTradeNo string) (state int, err error) {
	state, err = wxPayOrderQuery(ctx, outTradeNo)
	if state == 1 {
		err = model.NewOrder(vars.DBMysql).ActionOrder(outTradeNo)
	}
	return
}

func buildLocalOrder(mobile string) (*model.Order, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	order := &model.Order{
		Amt:        vars.YmlConfig.GetInt("WxPay.Amount"),
		Mobile:     mobile,
		OutTradeNo: node.Generate().String(),
		State:      vars.OrderStateCreated,
	}
	err = model.NewOrder(vars.DBMysql).CreateOrder(order)
	return order, err
}

// 生成一个长度为 NonceLength 的随机字符串（只包含大小写字母与数字）
func generateNonce() (string, error) {
	bytes := make([]byte, NonceLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	symbolsByteLength := byte(len(NonceSymbols))
	for i, b := range bytes {
		bytes[i] = NonceSymbols[b%symbolsByteLength]
	}
	return string(bytes), nil
}

func wxPayOrder(ctx *gin.Context, outTradeNo, ip, openid string, amt int64) (payUrl string, err error) {
	var (
		mchID                      = vars.YmlConfig.GetString("WxPay.Mchid")    // 商户号
		mchCertificateSerialNumber = vars.YmlConfig.GetString("WxPay.SerialNo") // 商户证书序列号
		mchAPIv3Key                = vars.YmlConfig.GetString("WxPay.ApiV3")    // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils2.LoadPrivateKeyWithPath(vars.BasePath + "/config/apiclient_key.pem")
	if err != nil {
		return "", fmt.Errorf("商户密钥加载失败: %s", err.Error())
	}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return "", fmt.Errorf("创建支付客户端失败: %s", err.Error())
	}

	svc := h5.H5ApiService{Client: client}
	//svc := jsapi.JsapiApiService{Client: client}
	resp, result, err := svc.Prepay(ctx,
		h5.PrepayRequest{
			Appid:       core.String(vars.YmlConfig.GetString("WxPay.AppId")),
			Mchid:       core.String(vars.YmlConfig.GetString("WxPay.Mchid")),
			Description: core.String(vars.YmlConfig.GetString("WxPay.Desc")),
			OutTradeNo:  core.String(outTradeNo),
			NotifyUrl:   core.String(vars.YmlConfig.GetString("WxPay.NotifyUrl")),
			Amount: &h5.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(amt),
			},
			//Payer: &jsapi.Payer{
			//	Openid: core.String(openid),
			//},
			SceneInfo: &h5.SceneInfo{
				PayerClientIp: core.String(ip),
				H5Info: &h5.H5Info{
					Type: core.String("Wap"),
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("调用支付失败: %s", err.Error())
	} else {
		// 处理返回结果
		if result.Response.StatusCode == 200 {
			fmt.Println("h5 下单完成：", resp.H5Url)
			payUrl = *resp.H5Url
			return
		} else {
			fmt.Println("调用支付失败", result.Response)
			return "", fmt.Errorf("调用支付失败: %d", result.Response.StatusCode)
		}
	}
}

func wxPayJsApiOrder(ctx *gin.Context, outTradeNo, ip, openid string, amt int64) (prepayId string, err error) {
	var (
		mchID                      = vars.YmlConfig.GetString("WxPay.Mchid")    // 商户号
		mchCertificateSerialNumber = vars.YmlConfig.GetString("WxPay.SerialNo") // 商户证书序列号
		mchAPIv3Key                = vars.YmlConfig.GetString("WxPay.ApiV3")    // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils2.LoadPrivateKeyWithPath(vars.BasePath + "/config/apiclient_key.pem")
	if err != nil {
		return "", fmt.Errorf("商户密钥加载失败: %s", err.Error())
	}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return "", fmt.Errorf("创建支付客户端失败: %s", err.Error())
	}

	svc := jsapi.JsapiApiService{Client: client}
	resp, result, err := svc.Prepay(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(vars.YmlConfig.GetString("WxPay.AppId")),
			Mchid:       core.String(vars.YmlConfig.GetString("WxPay.Mchid")),
			Description: core.String(vars.YmlConfig.GetString("WxPay.Desc")),
			OutTradeNo:  core.String(outTradeNo),
			NotifyUrl:   core.String(vars.YmlConfig.GetString("WxPay.NotifyUrl")),
			Amount: &jsapi.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(amt),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(openid),
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("调用支付失败: %s", err.Error())
	} else {
		// 处理返回结果
		if result.Response.StatusCode == 200 {
			fmt.Println("jsapi 下单完成：", resp.PrepayId)
			prepayId = *resp.PrepayId
			return
		} else {
			fmt.Println("调用支付失败", result.Response)
			return "", fmt.Errorf("调用支付失败: %d", result.Response.StatusCode)
		}
	}
}

// ActionOrder 异步回调 TODO
func ActionOrder(params *v_data.VApiWxPayAction) (err error) {
	// 非严谨支付流程，无需查询订单准确状态
	return model.NewOrder(vars.DBMysql).ActionOrder("")
}

func wxPayOrderQuery(ctx *gin.Context, outTradeNo string) (state int, err error) {
	var (
		mchID                      = vars.YmlConfig.GetString("WxPay.Mchid")    // 商户号
		mchCertificateSerialNumber = vars.YmlConfig.GetString("WxPay.SerialNo") // 商户证书序列号
		mchAPIv3Key                = vars.YmlConfig.GetString("WxPay.ApiV3")    // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils2.LoadPrivateKeyWithPath(vars.BasePath + "/config/apiclient_key.pem")
	if err != nil {
		return 0, fmt.Errorf("商户密钥加载失败: %s", err.Error())
	}
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return 0, fmt.Errorf("创建查询客户端失败: %s", err.Error())
	}

	svc := h5.H5ApiService{Client: client}
	//svc := jsapi.JsapiApiService{Client: client}
	resp, result, err := svc.QueryOrderByOutTradeNo(ctx,
		h5.QueryOrderByOutTradeNoRequest{
			Mchid:      core.String(vars.YmlConfig.GetString("WxPay.Mchid")),
			OutTradeNo: core.String(outTradeNo),
		},
	)

	if err != nil {
		return 0, fmt.Errorf("调用支付失败: %s", err.Error())
	} else {
		// 处理返回结果
		if result.Response.StatusCode == 200 {
			fmt.Println("查询完成：", resp)
			if *resp.TradeState == "SUCCESS" {
				return 1, nil
			}
			return 2, nil
		} else {
			fmt.Println("查询失败", result.Response)
			return 0, fmt.Errorf("查询失败: %d", result.Response.StatusCode)
		}
	}
}
