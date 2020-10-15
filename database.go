package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type params struct {
	dbms     string
	username string
	password string
	dbName   string
	dbHost   string
	dbPort   string
}

type users struct {
	Id       int `gorm:"primary_key,uniqueIndex"`
	Email    string
	Password string
}

// Functionr Create Tables and test user.
func initDB() {

	var config params

	// valuesEnv := godotenv.Load()
	// if valuesEnv != nil {
	// 	log.Fatal(valuesEnv)
	// }

	config.dbms = "mysql"       //os.Getenv("DBMS")  support  PostgreSQL - SQLite SQLServer
	config.username = "db_user" //os.Getenv("db_user")
	config.password = "db_pass" //os.Getenv("db_pass")
	config.dbName = "db_name"   // os.Getenv("db_name")
	config.dbHost = "db_host"   // os.Getenv("db_host")
	config.dbPort = "db_port"   //  os.Getenv("DB_PORT"),
	conn, err := gorm.Open(config.dbms, config.username+":"+config.password+"@("+config.dbHost+":"+config.dbPort+")/"+config.dbName+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db = conn
	err = db.Debug().CreateTable(Headset{}).Error
	if err != nil {
		log.Println("failed to create Headset Table")
	}
	err = db.Debug().CreateTable(Agent{}).Error
	if err != nil {
		log.Println("failed to create Agent Table")
	}
	err = db.Debug().CreateTable(users{}).AddUniqueIndex("email", "email").Error
	if err != nil {
		log.Println("failed to create Users Table")
	}
	err = db.Debug().Create(&users{
		Email:    "test@test.com",
		Password: "123456",
	}).Error
	if err != nil {
		log.Println(err)
	}
	err = db.Debug().CreateTable(users{}).AddUniqueIndex("email", "email").Error
	if err != nil {
		log.Println("failed to create Users Table")
	}
}

func (h *Headset) insertHeadset() {
	err := db.Debug().Create(&h).Error
	if err != nil {
		log.Println(err)
	}
}

func findAllHeadsets() []Headset {
	var report []Headset
	err := db.Debug().Find(&report).Error
	if err != nil {
		log.Println(err)
	}
	return report
}

func (h *Headset) assingHeadsetToAgent() {
	err := db.Debug().Model(h).Update(h).Error
	if err != nil {
		log.Println(err)
	}
}
