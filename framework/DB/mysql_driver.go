package DB

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"xorm.io/core"
)

type Config struct {
	Driver   string
	User     string
	Password string
	Host     string
	Name     string
	Prefix   string
	Charset  string
}

var (
	//定义orm引擎
	Engine *xorm.Engine
)

//xorm.exe reverse mysql "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8" templates/goxorm
var MasterDataSourceName string = "%s:%s@tcp(%s)/%s?charset=%s"

func Setup(cfg Config) error {
	var err error
	Engine, err = xorm.NewEngine(cfg.Driver, fmt.Sprintf(MasterDataSourceName,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Name,
		cfg.Charset))
	if err != nil {
		log.Fatal("connect mysql failure :", err)
		return err
	}
	//统一设置表前缀
	if cfg.Prefix != "" {
		tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, cfg.Prefix)
		Engine.SetTableMapper(tbMapper)
	}
	//Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	Engine.ShowSQL(true)
	//同步数据表结构，可以根据MODEL自动创建表，同时可以指定创建多个表
	//engine.Sync(new(UserInfo))
	return nil
}
