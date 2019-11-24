package models

type User struct {
	Id     int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name   string `json:"name" xorm:"not null default '' VARCHAR(20)"`
	Mobile string `json:"mobile" xorm:"not null default '' VARCHAR(20)"`
}
