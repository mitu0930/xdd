package models

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	"gopkg.in/yaml.v2"
)

type Yaml struct {
	Containers         []Container
	Tasks              []Task
	Qrcode             string
	Master             string
	Mode               string
	Static             string
	Database           string
	QywxKey            string `yaml:"qywx_key"`
	Resident           string
	UserAgent          string `yaml:"user_agent"`
	Theme              string
	TelegramBotToken   string `yaml:"telegram_bot_token"`
	TelegramUserID     int    `yaml:"telegram_user_id"`
	QQID               int64  `yaml:"qquid"`
	QQGroupID          int64  `yaml:"qqgid"`
	DefaultPriority    int    `yaml:"default_priority"`
	NoGhproxy          bool   `yaml:"no_ghproxy"`
	QbotPublicMode     bool   `yaml:"qbot_public_mode"`
	DailyAssetPushCron string `yaml:"daily_asset_push_cron"`
	Version            string `yaml:"version"`
	Node               string
	Npm                string
	Python             string
	Pip                string
}

var Balance = "balance"
var Parallel = "parallel"
var GhProxy = "https://ghproxy.com/"
var Cdle = false

var Config Yaml

func initConfig() {
	if ExecPath == "/Users/cdle/Desktop/xdd" {
		Cdle = true
	}
	confDir := ExecPath + "/conf"
	if _, err := os.Stat(confDir); err != nil {
		os.MkdirAll(confDir, os.ModePerm)
	}
	for _, name := range []string{"app.conf", "config.yaml"} {
		f, err := os.OpenFile(ExecPath+"/conf/"+name, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			logs.Warn(err)
		}
		s, _ := ioutil.ReadAll(f)
		if len(s) == 0 {
			logs.Info("下载配置%s", name)
			r, err := httplib.Get(GhProxy + "https://raw.githubusercontent.com/cdle/xdd/main/conf/" + name).Response()
			if err == nil {
				io.Copy(f, r.Body)
			}
		}
		f.Close()
	}
	config := ExecPath + "/conf/config.yaml"
	if Cdle {
		config = ExecPath + "/conf/config_cdle.yaml"
	}
	content, err := ioutil.ReadFile(config)
	if err != nil {
		logs.Warn("解析config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &Config) != nil {
		logs.Warn("解析config.yaml出错: %v", err)
	}
	if Config.Master == "" {
		Config.Master = "xxxx"
	}
	if Config.Mode != Parallel {
		Config.Mode = Balance
	}
	if Config.Qrcode != "" {
		Config.Theme = Config.Qrcode
	}
	if Config.NoGhproxy {
		GhProxy = ""
	}
	if Config.Database == "" {
		Config.Database = ExecPath + "/.xdd.db"
	}
	if Config.Npm == "" {
		Config.Npm = "npm"
	}
	if Config.Node == "" {
		Config.Node = "node"
	}
	if Config.Python == "" {
		Config.Python = "python3"
	}
	if Config.Pip == "" {
		Config.Pip = "Pip3"
	}
}
