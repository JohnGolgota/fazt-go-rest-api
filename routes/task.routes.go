package routes

import (
	"encoding/json"
	"net/http"

	"github.com/JohnGolgota/fazt-go-rest-api/database"
	"github.com/JohnGolgota/fazt-go-rest-api/models"
	"github.com/gorilla/mux"
)

type TaskRoute struct{}

func (t TaskRoute) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks", CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", GetTaskHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/tasks/{id}", DeleteTaskHandler).Methods("DELETE")
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := database.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	database.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	database.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewDecoder(r.Body).Decode(&task)
	database.DB.Save(&task)
	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	database.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	database.DB.Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}
