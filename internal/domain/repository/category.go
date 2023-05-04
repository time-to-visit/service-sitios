package repository

import (
	"service-sites/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryCategory interface {
	RegisterCategory(category entity.Category) (*entity.Category, error)
	FindCategoryById(idCateogry int) (*entity.Category, error)
	FindCategory() (*[]entity.Category, error)
	UpdateCategory(category entity.Category) (*entity.Category, error)
	DeleteCategory(idCategory int) error
}

type repositoryCategory struct {
	db *gorm.DB
}

func NewRepositoryCategory(db *gorm.DB) IRepositoryCategory {
	return &repositoryCategory{
		db,
	}
}

func (r *repositoryCategory) RegisterCategory(category entity.Category) (*entity.Category, error) {
	err := r.db.Create(&category).Error
	return &category, err
}

func (r *repositoryCategory) FindCategory() (*[]entity.Category, error) {
	var categorys []entity.Category
	err := r.db.Preload("SubCategory").Find(&categorys).Error
	return &categorys, err
}

func (r *repositoryCategory) UpdateCategory(category entity.Category) (*entity.Category, error) {
	err := r.db.Save(&category).Error
	return &category, err
}

func (r *repositoryCategory) DeleteCategory(idCategory int) error {
	return r.db.Delete(entity.Category{}, idCategory).Error
}

func (r *repositoryCategory) FindCategoryById(idCateogry int) (*entity.Category, error) {
	var category *entity.Category
	err := r.db.First(&category, idCateogry).Error
	return category, err
}
