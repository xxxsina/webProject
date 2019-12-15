package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"webProject/app/config"
	"webProject/com_party/libraries/DB"
)

type Config struct {
	App   app.Config
	Appmodel DB.Config
	Appredis DB.Rconfig
}

var (
	cfgPath = flag.String("base", "../../../src/webProject/com_party/config/base.yml", "appbase file path")
	cfgDbPath = flag.String("database", "../../../src/webProject/com_party/config/database.yml", "db file path")
)

/**
编译一：go build -o webexe main.go
 */
func main() {
	cfg := Config{}
	//基础配置信息读取
	if baseConfig, err := ioutil.ReadFile(*cfgPath); err != nil {
		panic(err)
	} else {
		//解析读取的配置数据到cfg
		if err := yaml.Unmarshal(baseConfig, &cfg); err != nil {
			panic(err)
		}
	}
	//数据库配置文件读取
	if dbConfig, err := ioutil.ReadFile(*cfgDbPath); err != nil {
		panic(err)
	} else {
		if err := yaml.Unmarshal(dbConfig, &cfg); err != nil {
			panic(err)
		}
	}
	//初始化数据库
	if err := DB.Setup(cfg.Appmodel); err != nil {
		panic(err)
	}
	//初始化redis
	if err := DB.Rsetup(cfg.Appredis); err != nil {
		//panic(err)
	}
	//初始化应用
	app.Run(cfg.App)
	log.Panicln("Server exiting")
}