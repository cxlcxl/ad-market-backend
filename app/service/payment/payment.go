package servicepayment

import (
	"fmt"
	"market/app/model"
	"market/app/vars"
	"market/library/curl"
	"math/rand"
	"time"
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
}

// UserOrder 用户下单
func UserOrder(mobile, ip string) (*PayResponse, error) {
	order, err := buildLocalOrder(mobile)
	if err != nil {
		return nil, err
	}
	var data = &PayRequest{
		AppId:       vars.YmlConfig.GetString("WxPay.AppId"),
		Mchid:       vars.YmlConfig.GetString("WxPay.Mchid"),
		Description: vars.YmlConfig.GetString("WxPay.Desc"),
		OurTradeNo:  order.OutTradeNo,
		NotifyUrl:   vars.YmlConfig.GetString("WxPay.NotifyUrl"),
		Amount:      PayAmt{Total: order.Amt, Currency: "CNY"},
		SceneInfo:   Scene{PayerClientIp: ip, H5Info: H5Info{Type: "Wap"}},
	}
	orderUrl := vars.YmlConfig.GetString("WxPay.H5.Order")
	c, err := curl.New(orderUrl).Debug(true).ResBody(true).Post().JsonData(data)
	if err != nil {
		return nil, err
	}
	var res PayResponse
	err = c.Request(&res, curl.JsonHeader())
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func buildLocalOrder(mobile string) (*model.Order, error) {
	sn := time.Now().Format("20060102150405")
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(85000) + 5000
	order := &model.Order{
		Amt:        vars.YmlConfig.GetInt("WxPay.Amount"),
		Mobile:     mobile,
		OutTradeNo: fmt.Sprintf("%s%d", sn, r),
	}
	err := model.NewOrder(vars.DBMysql).CreateOrder(order)
	return order, err
}
