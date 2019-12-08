package helper

import (
	"fmt"
	"time"
)

var Now time.Time

func init() {
	Now = time.Now()
}

//获取时间格式 2016-01-01
func TimeFormatYMD() string {
	return fmt.Sprintf("%d-%.2d-%.2d",
		Now.Year(),
		Now.Month(),
		Now.Day())
}

//获取时间格式 2016-01-01 01:01:01
func TimeFormatYMD_HIS() string {
	return fmt.Sprintf("%d-%.2d-%.2d %.2d:%.2d%.2d",
		Now.Year(),
		Now.Month(),
		Now.Day(),
		Now.Hour(),
		Now.Minute(),
		Now.Second())
}
