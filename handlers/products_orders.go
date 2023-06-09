package handlers

import (
	"encoding/json"
	productsordersdto "indochat/dto/products_orders"
	dto "indochat/dto/result"
	"indochat/models"
	"indochat/repositories"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerProducts_Orders struct {
	products_ordersRepository repositories.Products_OrdersRepository
}

func HandlerProducts_Orders(products_ordersRepository repositories.Products_OrdersRepository) *handlerProducts_Orders {
	return &handlerProducts_Orders{products_ordersRepository}
}

func (h *handlerProducts_Orders) FindProducts_Orders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orders, err := h.products_ordersRepository.FindProducts_Orders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: (orders)}
	json.NewEncoder(w).Encode(response)
	// fmt.Println(products)
}

func (h *handlerProducts_Orders) FindProducts_OrderByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	idUser := int(userInfo["id"].(float64))

	transactions, err := h.products_ordersRepository.GetProducts_Orders(idUser)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res := dto.ErrorResult{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	res := dto.SuccessResult{
		Code: http.StatusOK,
		Data: transactions,
	}
	json.NewEncoder(w).Encode(res)
}

func (h *handlerProducts_Orders) CreateProducts_Orders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request productsordersdto.Request_Products_Orders
	json.NewDecoder(r.Body).Decode(&request)

	validation := validator.New()
	errValidation := validation.Struct(request)
	if errValidation != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Failed Validation Transaction"}
		json.NewEncoder(w).Encode(response)
		return
	}
	products_orders := models.Products_Orders{
		Orders_Id:   request.Orders_Id,
		Products_Id: request.Products_Id,
	}

	// store data
	data, err := h.products_ordersRepository.CreateProducts_Orders(products_orders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
