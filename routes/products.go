package routes

import (
	"indochat/handlers"
	"indochat/pkg/middleware"
	"indochat/pkg/mysql"
	"indochat/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	r.HandleFunc("/products", middleware.Auth(h.FindProducts)).Methods("GET")
	r.HandleFunc("/products/{id}", h.GetProduct).Methods("GET")
	r.HandleFunc("/products", middleware.Auth(middleware.UploadFile(h.CreateProduct, "image"))).Methods("POST")
}
