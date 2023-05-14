package ali_sms

import (
	"errors"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"market/app/model"
	"market/app/utils"
	"market/app/vars"
	"market/library/curl"
	"time"
)

/*CreateClient
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

/*CreateClientWithSTS
* 使用STS鉴权方式初始化账号Client，推荐此方式。本示例默认使用AK&SK方式。
* @param accessKeyId
* @param accessKeySecret
* @param securityToken
* @return Client
* @throws Exception
 */
func CreateClientWithSTS(accessKeyId *string, accessKeySecret *string, securityToken *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
		// 必填，您的 Security Token
		SecurityToken: securityToken,
		// 必填，表明使用 STS 方式
		Type: tea.String("sts"),
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func SendSms(mobile, code string) (_err error) {
	accessKeyId := vars.YmlConfig.GetString("Sms.AccessKeyId")
	accessKeySecret := vars.YmlConfig.GetString("Sms.AccessKeySecret")
	client, _err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return _err
	}

	sign := vars.YmlConfig.GetString("Sms.SignName")
	tpl := vars.YmlConfig.GetString("Sms.TemplateCode")
	codeMsg := fmt.Sprintf(`{"code": "%s"}`, code)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(mobile),
		SignName:      tea.String(sign),
		TemplateCode:  tea.String(tpl),
		TemplateParam: tea.String(codeMsg),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := client.SendSmsWithOptions(sendSmsRequest, runtime)
	if _err != nil {
		return _err
	}

	console.Log(util.ToJSONString(resp))
	return _err
}

var smsCacheKey = "ali:sms"

// BuildAndSend 生成验证码并发送
func BuildAndSend(mobile string) (code string, err error) {
	key := fmt.Sprintf("%s:%s", smsCacheKey, mobile)
	if vars.DBRedis.HasKey(key) {
		code = vars.DBRedis.GetString(key)
		return
	} else {
		code, err = utils.GenValidateCode(4)
		if err != nil {
			return
		}

		if err = SendSms(mobile, code); err == nil {
			vars.DBRedis.SetString(key, code, time.Second*(10*60))
			_ = model.NewAct(vars.DBMysql).AccountCreate(vars.AccountStateNoAuth, mobile)
			return code, nil
		} else {
			fmt.Println("sms：", err)
			return "", errors.New("验证码存储失败，请重新发送")
		}
	}
}

// ValidSmsCode 验证并存储信息
func ValidSmsCode(mobile, code, logIdUrl string) (err error) {
	if vars.YmlConfig.GetBool("Debug") {
		eventReport(logIdUrl)
		return nil
	}
	key := fmt.Sprintf("%s:%s", smsCacheKey, mobile)
	if _code := vars.DBRedis.GetString(key); code != _code {
		return errors.New("验证码错误")
	} else {
		eventReport(logIdUrl)
		_ = vars.DBRedis.ExpireTime(key, 1)
		return nil
	}
}

func eventReport(logIdUrl string) {
	data := map[string]interface{}{
		"token": vars.YmlConfig.GetString("BDOCPC.TOKEN"),
		"conversionTypes": []map[string]interface{}{
			{"logidUrl": logIdUrl, "newType": 3},
		},
	}
	c, err := curl.New(vars.YmlConfig.GetString("BDOCPC.EventUrl")).Post().JsonData(data)
	if err != nil {
		fmt.Println("请求组合失败：", err)
		return
	}
	var response interface{}
	err = c.Request(&response, curl.JsonHeader())
	if err != nil {
		fmt.Println("请求失败：", err)
		return
	}
	fmt.Println("上报响应：", response)
}
