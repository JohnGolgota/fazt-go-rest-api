package main

import (
	"net/http"

	appConfig "github.com/JohnGolgota/fazt-go-rest-api/config"
	"github.com/JohnGolgota/fazt-go-rest-api/database"
	"github.com/JohnGolgota/fazt-go-rest-api/routes"
	"github.com/gorilla/mux"
)

func main() {

	database.DBConnection()
	// database.DB.AutoMigrate(models.User{}, models.Task{})

	port := appConfig.DefineApiPort()
	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler).Methods("GET")

	// Users routes
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(port, router)
}
