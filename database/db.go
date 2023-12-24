package database

import (
	"log"

	appConfig "github.com/JohnGolgota/fazt-go-rest-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN string = appConfig.DefineDatabaseDSN()
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Database connected")
}
