package handlers

import (
	"encoding/json"
	ordersdto "indochat/dto/orders"
	dto "indochat/dto/result"
	"indochat/models"
	"indochat/repositories"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerOrders struct {
	OrdersRepository repositories.OrdersRepository
}

func HandlerOrders(OrdersRepository repositories.OrdersRepository) *handlerOrders {
	return &handlerOrders{OrdersRepository}
}

// for view all data
func (h *handlerOrders) FindOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orders, err := h.OrdersRepository.FindOrders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: (orders)}
	json.NewEncoder(w).Encode(response)
	// fmt.Println(products)
}

func (h *handlerOrders) FindOrderByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	idUser := int(userInfo["id"].(float64))

	transactions, err := h.OrdersRepository.GetOrders(idUser)
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

func (h *handlerOrders) CreateOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request ordersdto.Request_Orders
	json.NewDecoder(r.Body).Decode(&request)

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	request.Customers_Id = int(userInfo["id"].(float64))

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
	localTime := time.Now()
	formattedTime := localTime.Format("Monday, 02-Jan-06 15:04:05 MST")

	orders := models.Orders{
		Date:         formattedTime,
		Status:       request.Status,
		// Product:      []models.Products{},
		Customers_Id: request.Customers_Id,
	}

	// store data
	data, err := h.OrdersRepository.CreateOrders(orders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertOrdersResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertOrdersResponse(r models.Orders) ordersdto.Response_Orders {
	return ordersdto.Response_Orders{
		Id:      r.Id,
		Date:    r.Date,
		Product: r.Product,
		Customers_Id: r.Customers_Id,
	}
}
