package models

type Categories struct {
	Id          int    `json:"id" gorm:"type: int;PRIMARY_KEY"`
	Name        string `json:"name"`
	Products_Id int    `json:"products_id"`
	Products    Products
}

func (Categories) TableName() string {
	return "categories"
}
