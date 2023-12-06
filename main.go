package main

import (
	"log"
	"tv360/src/db"
	"tv360/src/handler"
	"tv360/src/models"
	"tv360/src/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config: ", err)
	}
	connect, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Println("can not connect to database")
	} else {
		log.Println(connect)
	}

	store := db.NewStore(connect)

	err = connect.AutoMigrate(&models.Content2{})
	if err != nil {
		log.Fatal("can not migration Content2: ", err)
	}
	err = connect.AutoMigrate(&models.ContentTV{})
	if err != nil {
		log.Fatal("can not migration ContentTV: ", err)
	}

	server, err := handler.NewServer(config, store)
	if err != nil {
		log.Println("can not create server")
	}

	server.Start()
}
