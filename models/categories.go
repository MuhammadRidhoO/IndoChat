package models

type Categories struct {
	Id          int    `json:"id" gorm:"type: int;PRIMARY_KEY"`
	Name        string `json:"name"`
}

func (Categories) TableName() string {
	return "categories"
}
