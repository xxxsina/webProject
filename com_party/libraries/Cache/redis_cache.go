package Cache

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"time"
	"webProject/com_party/libraries/DB"
)

var CacheTimeOut time.Duration = 3600

type DialCacheOption struct {
	f func(*dialCacheOptions)
}

type dialCacheOptions struct {
	timeOut time.Duration
}

func DialCacheTimeOut(d time.Duration) DialCacheOption {
	return DialCacheOption{func(options *dialCacheOptions) {
		options.timeOut = d
	}}
}

//set设置
func Set(key string, data interface{}, options ...DialCacheOption) (bool, error) {
	conn := DB.RedisConn.Get()
	defer conn.Close()

	do := dialCacheOptions{}

	for _, option := range options {
		option.f(&do)
	}

	value, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	reply, err := redis.Bool(conn.Do("SET", key, value))
	if do.timeOut != 0 {
		conn.Do("EXPIRE", key, do.timeOut)
	} else {
		conn.Do("EXPIRE", key, CacheTimeOut)
	}

	return reply, err
}

//键值是否存在
func Exists(key string) bool {
	conn := DB.RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

//get获取
func Get(key string) ([]byte, error) {
	conn := DB.RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

//删除
func Delete(key string) (bool, error) {
	conn := DB.RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

//批量删除
func LikeDeletes(key string) error {
	conn := DB.RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
