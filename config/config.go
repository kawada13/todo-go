package config

import (
	"log"

	"gopkg.in/go-ini/ini.v1"
)

// データベース設定
type ConfigList struct {
	Port      string
	SQLDriver string
	Dbname    string
	LogFile   string
}

// どこからでもこの構造体が呼び出せるようにグローバルに変数を定義

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		Dbname:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
