package routes

// import (
// 	"indochat/handlers"
// 	"indochat/pkg/middleware"
// 	"indochat/pkg/mysql"
// 	"indochat/repositories"

// 	"github.com/gorilla/mux"
// )

// func Products_OrdersRoutes(r *mux.Router) {
// 	products_ordersRepository := repositories.RepositoryProducts_Orders(mysql.DB)
// 	h := handlers.HandlerProducts_Orders(products_ordersRepository)

// 	r.HandleFunc("/products-orders", middleware.Auth(h.FindProducts_Orders)).Methods("GET")
// 	r.HandleFunc("/products-orders/{id}", h.FindProducts_OrderByUser).Methods("GET")
// 	r.HandleFunc("/products-orders", middleware.Auth(h.CreateProducts_Orders)).Methods("POST")
// }
