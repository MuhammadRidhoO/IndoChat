package models


type Orders struct {
	Id           int `json:"id" gorm:"type:int;PRIMARY_KEY"`
	Date         string
	Status       string
	Product      []Products `json:"products" gorm:"many2many:products_orders;"`
	Customers    Customers
	Customers_Id int
}

func (Orders) TableName() string {
	return "orders"
}
