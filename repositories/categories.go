package repositories

import (
	"indochat/models"

	"gorm.io/gorm"
)

type CategoriesRepository interface {
	FindCategories() ([]models.Categories, error)
	GetCategories(Id int) (models.Categories, error)
	CreateCategories(categories models.Categories) (models.Categories, error)
	DeleteCategories(categories models.Categories) (models.Categories, error)
}

func RepositoryCategories(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) FindCategories() ([]models.Categories, error) {
	var categories []models.Categories
	err := r.db.Preload("Products").Find(&categories).Error

	return categories, err
}

func (r *repositories) GetCategories(Id int) (models.Categories, error) {
	var categories models.Categories
	err := r.db.First(&categories, Id).Error

	return categories, err
}
func (r *repositories) CreateCategories(categories models.Categories) (models.Categories, error) {
	err := r.db.Create(&categories).Error

	return categories, err
}

func (r *repositories) DeleteCategories(categories models.Categories) (models.Categories, error) {
	err := r.db.Delete(&categories).Error

	return categories, err
}
