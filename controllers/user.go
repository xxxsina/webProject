package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"webProject/service"
)

func GetName(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("type change failure:", err)
	}
	r := service.User{ID: id}
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
	fmt.Println(c)
	fmt.Println(c.Request.FormValue("name"))
	name := c.Request.FormValue("name")
	mobile := c.Request.FormValue("mobile")
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