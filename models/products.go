package models

type Products struct {
	Id          int          `json:"id" gorm:"type:int;PRIMARY_KEY"`
	Name        string       `josn:"name"`
	Price       int          `json:"price"`
	Descraption string       `json:"descraption"`
	Image       string       `json:"image"`
	Orders      []Orders     `json:"orders" gorm:"many2many:products_orders;"`
	Categories  []Categories `json:"categories" gorm:"foreignKey:Products_Id"`
}

func (Products) TableName() string {
	return "products"
}
