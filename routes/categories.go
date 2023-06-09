package routes

import (
	"indochat/handlers"
	"indochat/pkg/middleware"
	"indochat/pkg/mysql"
	"indochat/repositories"

	"github.com/gorilla/mux"
)

func CategoriesRoutes(r *mux.Router) {
	categoriesRepositories := repositories.RepositoryCategories(mysql.DB)
	h := handlers.HandlerCategories(categoriesRepositories)

	r.HandleFunc("/categories", middleware.Auth(h.FindCategories)).Methods("GET")
	r.HandleFunc("/categories/{id}", h.GetCategories).Methods("GET")
	r.HandleFunc("/categories", middleware.Auth(middleware.UploadFile(h.CreateCategories, "image"))).Methods("POST")
}
