package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	ProductRoutes(r)
	AuthRoutes(r)
	CategoriesRoutes(r)
	OrdersRoutes(r)
	// Products_OrdersRoutes(r)
}
