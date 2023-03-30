package vars

const (
	ApiPrefix       = "/api"
	ConfigKeyPrefix = "_ad_market_config_"
	LoginUserKey    = "ad_market_login_user"
	DateTimeFormat  = "2006-01-02 15:04:05"
	DateFormat      = "2006-01-02"
)

const (
	ResponseCodeOk = iota
	ResponseCodeError
	ResponseCodeOvertime
	ResponseCodeDatabaseErr
	ResponseCodeValidFailed
	ResponseCodeUnauthorized
	ResponseCodeEmptyToken
	ResponseCodeTokenErr
	ResponseCodeTokenExpire
)

const (
	AccountStateNoAuth = iota + 1
	AccountStateNoPaid
	AccountStatePaid
	AccountStateAdded
)

const (
	OrderStateCreated = iota + 1
	OrderStatePaid
	OrderStateDestroy
)

var (
	ResponseMsg = map[int]string{
		ResponseCodeOk:           "OK",
		ResponseCodeError:        "请求失败",
		ResponseCodeOvertime:     "请求超时",
		ResponseCodeDatabaseErr:  "数据库查询失败",
		ResponseCodeValidFailed:  "数据验证失败",
		ResponseCodeUnauthorized: "Unauthorized:权限不足",
		ResponseCodeEmptyToken:   "缺少 TOKEN",
		ResponseCodeTokenErr:     "TOKEN 错误",
		ResponseCodeTokenExpire:  "TOKEN 过期",
	}
	// AccountState 账号状态
	AccountState = map[int]string{
		AccountStateNoAuth: "未短信认证",
		AccountStateNoPaid: "已认证未支付",
		AccountStatePaid:   "已认证已支付",
		AccountStateAdded:  "已加微信",
	}
	// OrderState 订单状态
	OrderState = map[int]string{
		OrderStateCreated: "已创建",
		OrderStatePaid:    "已支付",
		OrderStateDestroy: "已作废",
	}
)
