package repositories

import (
	"indochat/models"

	"gorm.io/gorm"
)

type OrdersRepository interface {
	FindOrders() ([]models.Orders, error)
	GetOrders(Id int) (models.Orders, error)
	CreateOrders(orders models.Orders) (models.Orders, error)
	DeleteOrders(orders models.Orders) (models.Orders, error)
}

func RepositoryOrders(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindOrders() ([]models.Orders, error) {
	var orders []models.Orders
	err := r.db.Preload("Product").Find(&orders).Error

	return orders, err
}

func (r *repositories) GetOrders(Id int) (models.Orders, error) {
	var orders models.Orders
	err := r.db.First(&orders, "id = ?", Id).Error

	return orders, err
}
func (r *repositories) CreateOrders(orders models.Orders) (models.Orders, error) {
	err := r.db.Create(&orders).Error

	return orders, err
}

func (r *repositories) DeleteOrders(orders models.Orders) (models.Orders, error) {
	err := r.db.Delete(&orders).Error

	return orders, err
}
