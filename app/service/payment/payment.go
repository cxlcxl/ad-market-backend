package servicepayment

import (
	"crypto/rand"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	utils2 "github.com/wechatpay-apiv3/wechatpay-go/utils"
	"market/app/model"
	"market/app/vars"
)

type PayRequest struct {
	AppId       string `json:"appid"`        // 公众号APPID
	Mchid       string `json:"mchid"`        // 直连商户号
	Description string `json:"description"`  // 商品名称
	OurTradeNo  string `json:"out_trade_no"` // 商户订单号
	NotifyUrl   string `json:"notify_url"`   // 通知地址 要求必须为https地址
	Amount      PayAmt `json:"amount"`       // 金额
	SceneInfo   Scene  `json:"scene_info"`   // 场景信息
}

type PayAmt struct {
	Total    int    `json:"total"`    // 订单总金额，单位为分。
	Currency string `json:"currency"` // CNY
}

type Scene struct {
	PayerClientIp string `json:"payer_client_ip"` // IP
	H5Info        H5Info `json:"h5_info"`
}

type H5Info struct {
	Type string `json:"type"` // iOS, Android, Wap
}

type PayResponse struct {
	H5Url   string     `json:"h5_url"`
	Code    string     `json:"code"`
	Message string     `json:"message"`
	Detail  PayResBody `json:"detail"`
}

type PayResBody struct {
	Location string `json:"location"`
	Value    string `json:"value"`
}

const (
	// NonceSymbols 随机字符串可用字符集
	NonceSymbols = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// NonceLength 随机字符串的长度
	NonceLength = 32
)

// UserOrder 用户下单
func UserOrder(ctx *gin.Context, mobile string) (h5Url string, err error) {
	order, err := buildLocalOrder(mobile)
	if err != nil {
		return "", err
	}
	// 微信提供的包

	h5Url, err = wxPayOrder(ctx, order.OutTradeNo, ctx.ClientIP(), order.OpenId, int64(order.Amt))
	fmt.Println(err)
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

// ActionOrder 异步回调 TODO
func ActionOrder(outTradeNo string) (err error) {
	// 非严谨支付流程，无需查询订单准确状态
	return model.NewOrder(vars.DBMysql).ActionOrder(outTradeNo)
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
