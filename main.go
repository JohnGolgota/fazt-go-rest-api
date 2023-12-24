package main

import (
	"github.com/JohnGolgota/fazt-go-rest-api/database"
	"github.com/JohnGolgota/fazt-go-rest-api/server"
)

func main() {

	database.DBConnection()
	// database.DB.AutoMigrate(models.User{}, models.Task{})

	server := server.NewServer()
	server.Run()
}
