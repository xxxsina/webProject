package DB

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"time"
	"webProject/com_party/helper"
)

type Rconfig struct {
	Network     string
	Addr        string
	Password    string
	Db          int
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	Timeout     time.Duration
}

var (
	RedisConn *redis.Pool
)

func Rsetup(cfg Rconfig) error {
	RedisConn = &redis.Pool{
		MaxIdle:     cfg.MaxIdle,     //最大空闲连接数
		MaxActive:   cfg.MaxActive,   //在给定时间内，允许分配的最大连接数（当为零时，没有限制）
		IdleTimeout: cfg.IdleTimeout, //在给定时间内将会保持空闲状态，若到达时间限制则关闭连接（当为零时，没有限制）
		//提供创建和配置应用程序连接的一个函数
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(
				cfg.Network,
				cfg.Addr,
				redis.DialPassword(cfg.Password),
				redis.DialDatabase(cfg.Db))
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
		//可选的应用程序检查健康功能
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return errors.New(helper.CodeText(helper.Code10010))
}
