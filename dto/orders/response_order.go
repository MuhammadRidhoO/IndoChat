package ordersdto

import (
	"indochat/models"
)

type Response_Orders struct {
	Id           int    `json:"id"`
	Date         string `json:"date"`
	Status       string `json:"status"`
	Product      []models.Products
	Customers_Id int
}
