package productsdto

type Request_Products struct {
	Name          string `json:"name" validate:"required"`
	Price         int    `json:"price" validate:"required"`
	Descraption   string `json:"descraption" validate:"required"`
	Image         string `json:"image" form:"image"`
	Categories_Id int    `json:"categories_id" form:"categories_id" validate:"required"`
	Orders_Id     int    `json:"orders_id" form:"orders_id" validate:"required"`
}
