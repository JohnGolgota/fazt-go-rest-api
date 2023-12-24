package routes

import (
	"encoding/json"
	"net/http"

	"github.com/JohnGolgota/fazt-go-rest-api/database"
	"github.com/JohnGolgota/fazt-go-rest-api/models"
	"github.com/gorilla/mux"
)

type UserRoute struct{}

func (u UserRoute) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/users", GetUsersHandler).Methods("GET")
	router.HandleFunc("/users", CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUserHandler).Methods("DELETE")
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := database.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	database.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	database.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	database.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	database.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	database.DB.Delete(&user)
	w.WriteHeader(http.StatusOK)
}
