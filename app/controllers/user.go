package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/wumansgy/goEncrypt"
	"net/http"
	"webProject/com_party/helper"
	"webProject/com_party/libraries/Cache"
	"webProject/com_party/middleware"
	"webProject/com_party/service"
)

//这里提供一个参数提交的绑定数据结构
type UserInfo struct {
	Id     int    `form:"id" json:"id"`
	Name   string `form:"name" json:"name" binding:"required"`
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
}

// 绑定模型获取验证错误的方法
func (r *UserInfo) GetError(errs validator.ValidationErrors) (int, string) {
	fmt.Println("******************")
	fmt.Println(errs)
	fmt.Println("******************")
	for _, e := range errs {
		if e.Field() == "Mobile" {
			switch e.Tag() {
			case "required":
				return 10000, "请输入手机号码"
			}
		}
	}
	return 0, "参数错误"
}

func GetRedisUserInfo(c *gin.Context) {
	val, _ := Cache.Get("mykey")
	var str string
	_ = json.Unmarshal(val, &str)
	fmt.Println("\r\n val ===> ", str)

	c.JSON(http.StatusOK, gin.H{
		"code":    helper.Code200,
		"massage": str,
	})
}

type Endata struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Mobile string `json:"mobile"`
	Content string `json:"content"`
	Time string `json:"time"`
}

//对称加密 AES的CBC模式
func EnCrypt(c *gin.Context) {
	//需要加密输出的map
	data := gin.H{
		"id":      1,
		"name":    "xxx",
		"mobile":  "1903773333",
		"content": "床前明月光，疑是地上霜，举头望明月，低头思故乡",
		"time":    "2019-12-17",
	}
	//先转成json格式
	datajson, err := json.Marshal(data)
	//加密
	encrypt, err := goEncrypt.AesCbcEncrypt(datajson, []byte("e828f8db433920ed32f37ac3f9c9200a"))
	if err != nil {
		fmt.Println(err)
	}
	//解密为二进制
	decrypt, err := goEncrypt.AesCbcDecrypt(encrypt, []byte("e828f8db433920ed32f37ac3f9c9200a"))
	if err != nil {
		fmt.Println(err)
	}
	//解密后还要转成map，demap用来存解密后的数据
	var demap map[string]interface{}
	_ = json.Unmarshal(decrypt, &demap)

	fmt.Println("\r\n")
	fmt.Println("================================")
	fmt.Println(c.Request)
	fmt.Println("\r\n")
	for k, v := range demap {
		c.Set(k, v)
		//fmt.Println("k => ", k)
		//fmt.Println("v => ", v)
		//fmt.Println("\r\n")
	}
	fmt.Println("**********")
	fmt.Println(c.Request.GetBody)
	fmt.Println("**********")
	//fmt.Println(c)
	//fmt.Println(c.MustGet("name"))
	var endata Endata
	if err := c.Bind(&endata); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(c)
	fmt.Println("================================")
	fmt.Println("\r\n")

	c.JSON(http.StatusOK, gin.H{
		"code":    helper.Code200,
		"massage": "success",
		"encrypt": encrypt,
		"decrypt": demap,
		"name": endata.Name,
	})
}

//非对称加密RSA 获取公钥和私钥
func GetRsaKey(c *gin.Context) {
	_ = goEncrypt.GetRsaKey()

	c.JSON(http.StatusOK, nil)
}

func GetName(c *gin.Context) {
	//redis conn
	_, err := Cache.Set("mykey", "myval")
	e := Cache.Exists("mykey")
	b, er := Cache.Get("mykey")
	fmt.Println("vvvvvvvvvv")
	fmt.Println(err)
	fmt.Println(e)
	fmt.Println(b)
	fmt.Println(er)
	fmt.Println("^^^^^^^^^^^^")
	//ip := c.ClientIP()	//获取ip地址
	//fmt.Println("=============")
	//fmt.Println(ip)
	//fmt.Println("=============")
	//不绑定的获取提交的参数方式 http://127.0.0.1:8080/name/7
	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	log.Fatal("type change failure:", err)
	//}
	//r := service.User{ID: id}
	//h := md5.New()
	//h.Write([]byte("golang_bruce_lee"))
	//fmt.Println(hex.EncodeToString(h.Sum(nil)))
	//声明、绑定 http://127.0.0.1:8080/name/7?id=2
	var uinfo UserInfo
	if err := c.ShouldBind(&uinfo); err != nil {
		//code, msg := uinfo.GetError(err.(validator.ValidationErrors))
		fmt.Println(err)
		//fmt.Println(msg)
		//log.Fatal("bind user info failure :", code, msg)
		return
	}
	//绑定的获取参数方式
	r := service.User{ID: uinfo.Id}
	u, err := r.GetName()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//生成token
	token, err := middleware.JWTXcreate(u, c)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    helper.Code0,
			"massage": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    helper.Code200,
			"massage": helper.CodeText(helper.Code200),
			"token":   token,
			"data":    u,
		})
	}
}

func Add(c *gin.Context) {
	var name, mobile string
	fmt.Println("==================")
	//var v middleware.CustomClaims
	//v, _ = c.Get("claims")[0]
	//fmt.Println()
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	fmt.Println(claims.ID)
	fmt.Println("==================")
	//name = c.Request.FormValue("name")
	//mobile = c.Request.FormValue("mobile")
	var uinfo UserInfo
	if err := c.ShouldBindJSON(&uinfo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    helper.Code10003,
			"message": helper.CodeText(helper.Code10003),
			"data":    nil,
		})
		return
	} else {
		name = uinfo.Name
		mobile = uinfo.Mobile
	}
	r := service.User{Name: name, Mobile: mobile}
	u, err := r.InsertUser()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    helper.Code200,
		"massage": helper.CodeText(helper.Code200),
		"data":    u,
	})
}
