package repository

import (
	"service-sites/internal/domain/entity"
	"service-sites/internal/domain/utils"

	"gorm.io/gorm"
)

type IRepositorySites interface {
	InsertSites(review entity.Sites) (*entity.Sites, error)
	FindOneSites(idSites int) (*entity.Sites, error)
	DeleteSites(idSites int) error
	FindSites(filter map[string]interface{}) (*[]entity.Sites, error)
}

type repositorySites struct {
	db *gorm.DB
}

func NewRepositorySites(db *gorm.DB) IRepositorySites {
	return &repositorySites{
		db,
	}
}

func (r *repositorySites) InsertSites(review entity.Sites) (*entity.Sites, error) {
	err := r.db.Create(&review).Error
	return &review, err
}
func (r *repositorySites) FindOneSites(idSites int) (*entity.Sites, error) {
	var sites entity.Sites
	err := r.db.Preload("Contact").Preload("Resource").Preload("Category").First(&sites, idSites).Error
	return &sites, err
}

func (r *repositorySites) DeleteSites(idSites int) error {
	return r.db.Delete(entity.Sites{}, idSites).Error
}
func (r *repositorySites) FindSites(filter map[string]interface{}) (*[]entity.Sites, error) {
	var sites []entity.Sites
	command, request := utils.GetWhere(filter)
	err := r.db.Preload("Contact").
		Preload("Resource").
		Preload("Category").
		Where(command, request...).Find(&sites).Error
	return &sites, err
}
