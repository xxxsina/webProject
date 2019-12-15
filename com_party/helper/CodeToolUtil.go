package helper

const (
	Code0				= 0
	Code200				= 200
	Code10000           = 9998 + iota	//iota是根据当前位置来的 比如现在这个位置从2开始计数
	Code10001
	Code10002
	Code10003
	Code10004
	Code10005
	Code10006
	Code10007
	Code10008
	Code10009
	Code10010
)

var codeText = map[int]string{
	Code0:           	 "操作失败",
	Code200:           	 "操作成功",
	Code10000:           "登录失败",
	Code10001:           "密码错误",
	Code10002:           "登录过期",
	Code10003:           "数据获取失败",
	Code10004:           "请求未携带令牌",
	Code10005:           "授权已过期",
	Code10006:           "令牌尚未激活",
	Code10007:           "令牌错误",
	Code10008:           "令牌解析失败",
	Code10009:           "令牌生成失败",
	Code10010:           "连接redis失败",
}

func CodeText(code int) string {
	return codeText[code]
}