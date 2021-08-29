package models

import (
	"strings"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/httplib"
)

var ua = "'jdapp;android;10.0.5;11;0393465333165363-5333430323261366;network/wifi;model/M2102K1C;osVer/30;appBuild/88681;partner/lc001;eufv/1;jdSupportDarkMode/0;Mozilla/5.0 (Linux; Android 11; M2102K1C Build/RKQ1.201112.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/77.0.3865.120 MQQBrowser/6.2 TBS/045534 Mobile Safari/537.36'"

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
