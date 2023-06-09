package handlers

import (
	"encoding/json"
	"fmt"
	productsdto "indochat/dto/products"
	dto "indochat/dto/result"
	"indochat/models"
	"indochat/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerProduct struct {
	ProductRepository repositories.ProductsRepository
}

func HandlerProduct(ProductRepository repositories.ProductsRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

// for view all data
func (h *handlerProduct) FindProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	product, err := h.ProductRepository.FindProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: (product)}
	json.NewEncoder(w).Encode(response)
	// fmt.Println(products)
}

func (h *handlerProduct) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//params
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//get data
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	//to view success get data
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: product}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	orders_id, _ := strconv.Atoi(r.FormValue("orders_id"))
	categories_id, _ := strconv.Atoi(r.FormValue("categories_id"))

	request := productsdto.Request_Products{
		Name:          r.FormValue("name"),
		Image:         r.FormValue("image"),
		Price:         price,
		Categories_Id: categories_id,
		Descraption:   r.FormValue("descraption"),
		Orders_Id:     orders_id,
	}

	// validation
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	bill := int(request.Price)
	discount := int(10)
	afterDiscount := int(bill - (bill * discount / 100))

	product := models.Products{
		Name:          request.Name,
		Descraption:   request.Descraption,
		Price:         afterDiscount,
		Categories_Id: request.Categories_Id,
		Orders_Id:     request.Orders_Id,
		Image:         filename,
	}

	// store data
	data, err := h.ProductRepository.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	productGet, err := h.ProductRepository.GetProduct(data.Id)
	productGet.Image = os.Getenv("PATH_FILE") + productGet.Image
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// success
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
