package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
	"webProject/com_party/helper"
)

const FILE_SUFFIX = ".log"

// 日志记录到文件
func LoggerToFile(logFilePath string, logFileName string, debugLevel uint32) gin.HandlerFunc {
	// 日志文件
	//fileName := path.Join(filePath, logFileName)
	var fileAllPathName string

	//设置带时间的文件和文件夹
	xTime := helper.TimeFormatYMD()
	fileName := logFileName + "-" + xTime + FILE_SUFFIX

	//检查时间文件夹是否存在，不存在就创建
	filePath := path.Join(logFilePath, xTime)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		err := os.Mkdir(filePath, os.ModePerm)
		if err == nil {
			//注意这里是有命名空间问题的，需要在外面先声明变量
			fileAllPathName = path.Join(filePath, fileName)
		} else {
			fileAllPathName = path.Join(logFilePath, fileName)
		}
	} else {
		fileAllPathName = path.Join(filePath, fileName)
	}

	// 写入文件
	src, err := os.OpenFile(fileAllPathName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println("err : ", err)
	}

	// 实例化
	logger := logrus.New()

	// 设置输出
	logger.Out = src

	// 设置日志级别
	//logger.SetLevel(logrus.DebugLevel)
	logger.SetLevel(logrus.Level(debugLevel))

	// 设置 rotatelogs
	//logWriter, err := rotatelogs.New(
	//	// 分割后的文件名称
	//	fileName + ".%Y%m%d.log",
	//
	//	// 生成软链，指向最新日志文件
	//	rotatelogs.WithLinkName(fileName),
	//
	//	// 设置最大保存时间(7天)
	//	rotatelogs.WithMaxAge(7*24*time.Hour),
	//
	//	// 设置日志切割时间间隔(1天)
	//	rotatelogs.WithRotationTime(24*time.Hour),
	//)

	//writeMap := lfshook.WriterMap{
	//	logrus.InfoLevel:  logWriter,
	//	logrus.FatalLevel: logWriter,
	//	logrus.DebugLevel: logWriter,
	//	logrus.WarnLevel:  logWriter,
	//	logrus.ErrorLevel: logWriter,
	//	logrus.PanicLevel: logWriter,
	//}

	//lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
	//	TimestampFormat:"2006-01-02 15:04:05",
	//})

	// 新增 Hook
	//logger.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 换一下日期格式
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat:"2006-01-02 15:04:05",
		})

		// 日志格式
		fields := make(logrus.Fields)
		fields["status_code"] 	= statusCode
		fields["latency_time"] 	= latencyTime
		fields["client_ip"] 	= clientIP
		fields["req_method"] 	= reqMethod
		fields["req_uri"] 		= reqUri
		if reqMethod == "POST" {
			fields["params"] = c.Request.Body
		}
		logger.WithFields(fields).Info()

		// 处理请求，要放在最后，要不然params获取不到
		c.Next()
	}
}

// 日志记录到 MongoDB
//func LoggerToMongo() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//	}
//}
//
//// 日志记录到 ES
//func LoggerToES() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//	}
//}
//
//// 日志记录到 MQ
//func LoggerToMQ() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//	}
//}