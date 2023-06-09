package models

type Orders struct {
	Id           int `json:"id" gorm:"type:int;PRIMARY_KEY"`
	Date         string
	Status       string
	Products     []Products `json:"products" gorm:"foreignKey:Orders_Id"`
	Customers    Customers
	Customers_Id int
}

func (Orders) TableName() string {
	return "orders"
}
