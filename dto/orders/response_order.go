package ordersdto

import (
	"indochat/models"
)

type Response_Orders struct {
	Id           int    `json:"id"`
	Date         string `json:"date"`
	Status       string `json:"status"`
	Products     []models.Products
	Products_Id  int
	Customers_Id int
}
