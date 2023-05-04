package repository

import (
	"service-sites/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositorySubCategory interface {
	InsertSubcategory(review entity.Subcategory) (*entity.Subcategory, error)
	FindOneSubcategory(idSubCategory int) (*entity.Subcategory, error)
	DeleteSubcategory(idCategory int, idSubcategory int) error
	FindSubcategory(idCategory int) (*[]entity.Subcategory, error)
}

type repositorySubCategory struct {
	db *gorm.DB
}

func NewRepositorySubCategory(db *gorm.DB) IRepositorySubCategory {
	return &repositorySubCategory{
		db,
	}
}

func (r *repositorySubCategory) InsertSubcategory(subcategory entity.Subcategory) (*entity.Subcategory, error) {
	err := r.db.Create(&subcategory).Error
	return &subcategory, err
}
func (r *repositorySubCategory) FindOneSubcategory(idSubCategory int) (*entity.Subcategory, error) {
	var subcategory entity.Subcategory
	err := r.db.First(&subcategory, idSubCategory).Error
	return &subcategory, err
}
func (r *repositorySubCategory) DeleteSubcategory(idCategory int, idSubcategory int) error {
	err := r.db.Where("category_id = ?  and id= ? ", idCategory, idSubcategory).Delete(entity.Subcategory{}).Error
	return err
}
func (r *repositorySubCategory) FindSubcategory(idCategory int) (*[]entity.Subcategory, error) {
	var subcategorys []entity.Subcategory
	err := r.db.Where("category_id = ?", idCategory).Find(&subcategorys).Error
	return &subcategorys, err
}
