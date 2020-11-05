package main

import (
	"fmt"
	"./Config"
	"./Models"
	"./Controllers"
	"github.com/jinzhu/gorm"
)

var err error

func main(){
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()

	r.Run()
}