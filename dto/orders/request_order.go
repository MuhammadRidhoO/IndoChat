package ordersdto

import (
	"indochat/models"
)

type Request_Orders struct {
	Date         string `json:"date"`
	Status       string `json:"status"`
	Product      models.Products
	Customers_Id int
}
