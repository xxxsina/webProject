package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"webProject/service"
)

//这里提供一个参数提交的绑定数据结构
type UserInfo struct {
	Id int `form:"id" json:"id"`
	Name string `form:"name" json:"name" binding:"required"`
	Mobile string `form:"name" json:"mobile" binding:"required"`
}

func GetName(c *gin.Context) {
	//不绑定的获取提交的参数方式 http://127.0.0.1:8080/name/7
	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	log.Fatal("type change failure:", err)
	//}
	//r := service.User{ID: id}
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

	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"massage" : "success",
		"data" : u,
	})
}

func Add(c *gin.Context)  {
	var name, mobile string
	//name = c.Request.FormValue("name")
	//mobile = c.Request.FormValue("mobile")
	var uinfo UserInfo
	if err := c.ShouldBindJSON(&uinfo); err != nil {
		fmt.Println("我没绑定:", err)
		c.JSON(http.StatusOK, gin.H{
			"code" : 0,
			"message" : "获取数据失败",
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
		"code" : http.StatusOK,
		"massage" : "success",
		"data" : u,
	})
}