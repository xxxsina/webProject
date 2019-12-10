package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"webProject/com_party/helper"
	"webProject/com_party/middleware"
	"webProject/com_party/service"
)

//这里提供一个参数提交的绑定数据结构
type UserInfo struct {
	Id int `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Mobile string `form:"name" json:"mobile"`
}

func GetName(c *gin.Context) {
	ip := c.ClientIP()//获取ip地址
	fmt.Println("=============")
	fmt.Println(ip)
	fmt.Println("=============")
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
		log.Fatal("bind user info failure :", err)
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
			"code" : helper.Code0,
			"massage" : err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code" : helper.Code200,
			"massage" : helper.CodeText(helper.Code200),
			"token" : token,
			"data" : u,
		})
	}
}

func Add(c *gin.Context)  {
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
			"code" : helper.Code10003,
			"message" : helper.CodeText(helper.Code10003),
			"data" : nil,
		})
		return
	} else {
		name = uinfo.Name
		mobile = uinfo.Mobile
	}
	r := service.User{Name:name, Mobile: mobile}
	u, err := r.InsertUser()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : helper.Code200,
		"massage" : helper.CodeText(helper.Code200),
		"data" : u,
	})
}