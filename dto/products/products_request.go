package productsdto

import "indochat/models"

type Request_Products struct {
	Name        string `json:"name" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Descraption string `json:"descraption" validate:"required"`
	Image       string `json:"image" form:"image"`
	Orders      models.Orders
	Categories  models.Categories
}
