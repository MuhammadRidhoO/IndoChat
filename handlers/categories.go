package handlers

import (
	"encoding/json"
	"fmt"
	categoriestdto "indochat/dto/categories"
	dto "indochat/dto/result"
	"indochat/models"
	"indochat/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerCategories struct {
	CategoriesRepository repositories.CategoriesRepository
}

func HandlerCategories(CategoriesRepository repositories.CategoriesRepository) *handlerCategories {
	return &handlerCategories{CategoriesRepository}
}

func (h *handlerCategories) FindCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories, err := h.CategoriesRepository.FindCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: (categories)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCategories) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//params
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//get data
	categories, err := h.CategoriesRepository.GetCategories(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	//to view success get data
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: categories}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerCategories) CreateCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := categoriestdto.Request_Categories{
		Name:        r.FormValue("name"),
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
	categories := models.Categories{
		Name:        request.Name,
	}

	// store data
	data, err := h.CategoriesRepository.CreateCategories(categories)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// success
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertCategoriesResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertCategoriesResponse(r models.Categories) categoriestdto.Response_Categories {
	return categoriestdto.Response_Categories{
		Id:          r.Id,
		Name:        r.Name,
	}
}
