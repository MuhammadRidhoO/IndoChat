package models

type Customers struct {
	Id              int      `json:"id"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Password        string   `json:"password"`
}

func (Customers) TableName() string {
	return "customers"
}
