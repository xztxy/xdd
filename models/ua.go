package models

import (
	"strings"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/httplib"
)

var ua = "Mozilla/5.0 (iPhone; CPU iPhone OS 14_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/92.0.4515.90 Mobile/15E148 Safari/604.1"

func initUserAgent() {
	u := &UserAgent{}
	err := db.Order("id desc").First(u).Error
	if err != nil && strings.Contains(err.Error(), "converting") {
		db.Migrator().DropTable(&UserAgent{})
		Daemon()
	}
	if u.Content != "" {
		ua = u.Content
	} else {
		if Config.UserAgent != "" {
			logs.Info("使用自定义User-Agent")
			ua = Config.UserAgent
		} else {
			logs.Info("更新User-Agent")
			var err error
			ua, err = httplib.Get(GhProxy + "https://raw.githubusercontent.com/xztxy/xdd/main/ua.txt").String()
			if err != nil {
				logs.Info("更新User-Agent失败")
			}
		}
	}
}

func GetUserAgent() string {
	return ua
}

type UserAgent struct {
	ID      int
	Content string
}
