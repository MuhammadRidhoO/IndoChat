package routes

import (
	"indochat/handlers"
	"indochat/pkg/middleware"
	"indochat/pkg/mysql"
	"indochat/repositories"

	"github.com/gorilla/mux"
)

func OrdersRoutes(r *mux.Router) {
	OrdersRepositories := repositories.RepositoryOrders(mysql.DB)
	h := handlers.HandlerOrders(OrdersRepositories)

	r.HandleFunc("/orders", middleware.Auth(h.FindOrders)).Methods("GET")
	r.HandleFunc("/orders/{id}", middleware.Auth(h.FindOrderByUser)).Methods("GET")
	r.HandleFunc("/orders", middleware.Auth(h.CreateOrders)).Methods("POST")
}
