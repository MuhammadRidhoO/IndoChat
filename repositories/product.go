package repositories

import (
	"indochat/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProducts() ([]models.Products, error)
	GetProduct(Id int) (models.Products, error)
	CreateProduct(product models.Products) (models.Products, error)
	DeleteProduct(Product models.Products) (models.Products, error)
}

func RepositoryProduct(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindProducts() ([]models.Products, error) {
	var product []models.Products
	err := r.db.Preload("Orders").Preload("Categories").Find(&product).Error

	return product, err
}

func (r *repositories) GetProduct(Id int) (models.Products, error) {
	var product models.Products
	err := r.db.Preload("Orders").Preload("Categories").First(&product, Id).Error

	return product, err
}
func (r *repositories) CreateProduct(product models.Products) (models.Products, error) {
	err := r.db.Create(&product).Error

	return product, err
}

func (r *repositories) DeleteProduct(product models.Products) (models.Products, error) {
	err := r.db.Delete(&product).Error

	return product, err
}
