package repositories

import (
	"indochat/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(Customers models.Customers) (models.Customers, error)
	Login(email string) (models.Customers, error)
	CheckOut(Id int) (models.Customers, error)
}

func RepositoryAuth(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) Register(customers models.Customers) (models.Customers, error) {
	err := r.db.Create(&customers).Error

	return customers, err
}

func (r *repositories) Login(email string) (models.Customers, error) {
	var customers models.Customers
	err := r.db.First(&customers, "email=?", email).Error

	return customers, err
}

func (r *repositories) CheckOut(Id int) (models.Customers, error) {
	var customer models.Customers
	err := r.db.First(&customer, Id).Error

	return customer, err
}
