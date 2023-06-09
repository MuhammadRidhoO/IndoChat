package models

type Products struct {
	Id            int        `json:"id" gorm:"type:int;PRIMARY_KEY"`
	Name          string     `josn:"name"`
	Price         int        `json:"price"`
	Descraption   string     `json:"descraption"`
	Image         string     `json:"image"`
	Orders        Orders     `json:"orders"`
	Orders_Id     int        `json:"orders_id"`
	Categories    Categories `json:"categories"`
	Categories_Id int        `json:"categories_id"`
}

func (Products) TableName() string {
	return "products"
}
