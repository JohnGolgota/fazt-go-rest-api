package server

import (
	"net/http"

	appConfig "github.com/JohnGolgota/fazt-go-rest-api/config"

	"github.com/JohnGolgota/fazt-go-rest-api/routes"
	"github.com/gorilla/mux"
)

type Server struct {
	// ...
}

func (s *Server) Run() {
	port := appConfig.DefineApiPort()
	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler).Methods("GET")

	// Users routes
	userRoute := routes.NewRouterService(&routes.UserRoute{})
	userRoute.SetupRoutes(router)

	// Tasks routes
	taskRoute := routes.NewRouterService(&routes.TaskRoute{})
	taskRoute.SetupRoutes(router)

	http.ListenAndServe(port, router)
}

func NewServer() *Server {
	return &Server{}
}
