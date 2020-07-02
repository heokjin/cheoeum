package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var (
	db *gorm.DB
)

func InitDB() {
	var err  error
	//db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 dbname=scott sslmode=disable")
	if err != nil {
		log.Println("Got error when connect database", err)
	}
	log.Println("DB OK:", db.DB())
	log.Println("DB OK:", db.Dialect().GetName())


	// for task table
	db.LogMode(true)
}

func InitSchema() {
	db.AutoMigrate(
		&User{},
	)
}