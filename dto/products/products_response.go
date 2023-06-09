package productsdto

import "indochat/models"

type Response_Products struct {
	Id          int
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Descraption string `json:"descraption"`
	Image       string `json:"image"`
	Orders      models.Orders
	Categories  models.Categories
}
