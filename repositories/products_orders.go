package repositories

// import (
// 	"indochat/models"

// 	"gorm.io/gorm"
// )

// type Products_OrdersRepository interface {
// 	FindProducts_Orders() ([]models.Products_Orders, error)
// 	GetProducts_Orders(Id int) (models.Products_Orders, error)
// 	CreateProducts_Orders(Products_Orders models.Products_Orders) (models.Products_Orders, error)
// 	DeleteProducts_Orders(Products_Orders models.Products_Orders) (models.Products_Orders, error)
// }

// func RepositoryProducts_Orders(db *gorm.DB) *repositories {
// 	return &repositories{db}
// }

// func (r *repositories) FindProducts_Orders() ([]models.Products_Orders, error) {
// 	var product []models.Products_Orders
// 	err := r.db.Find(&product).Error

// 	return product, err
// }

// func (r *repositories) GetProducts_Orders(Id int) (models.Products_Orders, error) {
// 	var products_orders models.Products_Orders
// 	err := r.db.First(&products_orders, "id = ?", Id).Error

// 	return products_orders, err
// }
// func (r *repositories) CreateProducts_Orders(products_orders models.Products_Orders) (models.Products_Orders, error) {
// 	err := r.db.Create(&products_orders).Error

// 	return products_orders, err
// }

// func (r *repositories) DeleteProducts_Orders(products_orders models.Products_Orders) (models.Products_Orders, error) {
// 	err := r.db.Delete(&products_orders).Error

// 	return products_orders, err
// }
