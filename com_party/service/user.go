package service

import (
	"fmt"
	"log"
	"webProject/com_party/libraries/DB"
	"webProject/com_party/models"
)

type User struct {
	ID     int
	Name   string
	Mobile string
}

func (r *User) GetName() (*models.User, error) {
	u := &models.User{Id: r.ID}
	flag, err := DB.Engine.Get(u)
	if flag {
		fmt.Println("%s", r)
	} else if err != nil {
		log.Fatal("error", err)
	} else {
		fmt.Println(" id = ", r.ID)
	}
	return u, err
}

func (r *User) InsertUser() (*models.User, error) {
	u := &models.User{Name: r.Name, Mobile: r.Mobile}
	num, err := DB.Engine.InsertOne(u)
	if err != nil {
		log.Fatal("error : ", err)
	} else {
		fmt.Println("sync success count：", num)
		fmt.Println("%v\n", u)
	}
	return u, err
}
