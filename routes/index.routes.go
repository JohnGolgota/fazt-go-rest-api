package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route interface {
	SetupRoutes(router *mux.Router)
}

type RouteService struct {
	routes Route
}

func NewRouterService(rs Route) *RouteService {
	return &RouteService{routes: rs}
}

func (rs *RouteService) SetupRoutes(router *mux.Router) {
	rs.routes.SetupRoutes(router)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK melo"))
}
