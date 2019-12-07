package routers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"webProject/app/controllers"
	"webProject/com_party/middleware"
)

type Config struct {
	Port string			//监听端口
	Debug bool			//是否开启日志
	Debuglevel uint32	//日志级别
	Logfilepath string	//日志路径
	Logfilename string	//日志名称
}

//解决gin只能从Request.body获取一次的方法
func reWriteParamsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := c.GetRawData()
		if err == nil {
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))	//关键点
		}
		c.Next()
	}
}

func RegisterRouters(g *gin.Engine, cfg Config) {
	//全局中间件
	{
		//重新设置参数
		g.Use(reWriteParamsMiddleware())
		//Logger
		if cfg.Debug {
			g.Use(middleware.LoggerToFile(cfg.Logfilepath, cfg.Logfilename, cfg.Debuglevel))
		}
	}
	//路由
	{
		//单个路由方法
		g.GET("/name/:id", controllers.GetName)
		//一组路由方法
		v1 := g.Group("/v1", middleware.JWTAuth())
		{
			v1.POST("/add", controllers.Add)
		}
	}
}