package app

type Config struct {
	Port string			//监听端口
	Debug bool			//是否开启日志
	Debuglevel uint32	//日志级别
	Logfilepath string	//日志路径
	Logfilename string	//日志名称
	Jwtsignkey string	//JWT SignKey
}