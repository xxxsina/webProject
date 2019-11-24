package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"webProject/app"
	"webProject/models"
)

type Config struct {
	App app.Config
	Model models.Config
}

var (
	cfgPath = flag.String("config", "config/config.yml", "appbase file path")
)

/**
编译一：go build -o webexe main.go
 */
func main() {
	cfg := Config{}
	//读取配置文件到data
	if data, err := ioutil.ReadFile(*cfgPath); err != nil {
		panic(err)
	} else {
		//解析读取的配置数据到cfg
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			panic(err)
		}
	}
	//初始化数据库
	if err := models.Setup(cfg.Model); err != nil {
		panic(err)
	}
	//初始化应用
	app.Run(cfg.App)
	log.Panicln("Server exiting")
}