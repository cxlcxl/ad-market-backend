package vars

import (
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
	"market/library/config_interface"
	myRedis "market/library/redis"
)

var (
	BasePath  string
	DBMysql   *gorm.DB
	YmlConfig config_interface.YamlConfigInterface
	DBRedis   *myRedis.DBRedis
)

type LoginUser struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
}

func init() {
	if dir, err := os.Getwd(); err != nil {
		log.Fatal("文件目录获取失败")
		return
	} else {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(dir, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = dir
		}
	}
}
