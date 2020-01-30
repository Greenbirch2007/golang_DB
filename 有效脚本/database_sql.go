package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
	//log "github.com/code.google.com/log4go"
)


type Userinfo struct {
	Id int `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Phone string `gorm:"column:phone" json:phone`
	CreatedTime time.Time `gorm:"column:created_time" json:"create_time"`
	CreatedUser string `gorm:"column:created_user" json:"created_user"`
}


func main(){
	dbHost:="user:password@tcp(127.0.0.1:3306)/test?"+
	"charset=utf8&&parseTime=True&loc=Local"
	db,err := gorm.Open("mysql",dbHost)
	if err != nil{
		fmt.Println("数据库链接异常")
	}

	var userInfos []UserInfo
	db.LogMode(true) //打印sql语句
	err1 := db.Table("user_info").Find(&userInfos).Error
	if err1 !=nil{
		fmt.Println("查询出错%s",err1.Error())
	}

	for _,v := range userInfos{
		fmt.Printf("userInfos:%#v",v)
	}
}