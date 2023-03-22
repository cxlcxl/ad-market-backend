package vars

const (
	ApiPrefix = "/api"

	ConfigKeyPrefix = "_ad_market_config_"
	LoginUserKey    = "ad_market_login_user"

	UserStateValid = 1

	MaxPageSize       uint64 = 100
	SystemDefaultPass        = "a123456"

	DateTimeFormat = "2006-01-02 15:04:05"

	DateFormat = "2006-01-02"
	Env        = "dev"
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
	CommonStateVoid = iota
	CommonStateValid
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
	// CommonState 通用数据库状态字段
	CommonState = map[int]string{
		CommonStateVoid:  "停用",
		CommonStateValid: "正常",
	}
)
