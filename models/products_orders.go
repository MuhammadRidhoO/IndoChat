package models

type Products_Orders struct {
	Id          int `json:"id" gorm:"type:int;PRIMARY_KEY"`
	Orders_Id   int `json:"orders_id"`
	Products_Id int `json:"products_id"`
}

func (Products_Orders) TableName() string {
	return "products_orders"
}
