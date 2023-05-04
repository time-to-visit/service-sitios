package repository

import (
	"service-sites/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryResource interface {
	InsertResource(resource entity.Resource) (*entity.Resource, error)
	FindResource(idSites int) (*[]entity.Resource, error)
	FindOneResource(idSites int, idResource int) (*entity.Resource, error)
	DeleteResource(idSites int, idResource int) error
}

type repositoryResource struct {
	db *gorm.DB
}

func NewRepositoryResource(db *gorm.DB) IRepositoryResource {
	return &repositoryResource{
		db,
	}
}

func (r *repositoryResource) InsertResource(resource entity.Resource) (*entity.Resource, error) {
	err := r.db.Create(&resource).Error
	return &resource, err
}

func (r *repositoryResource) FindResource(idSites int) (*[]entity.Resource, error) {
	var resource []entity.Resource
	err := r.db.Where("sites_id = ?", idSites).Find(&resource).Error
	return &resource, err
}

func (r *repositoryResource) DeleteResource(idSites int, idResource int) error {
	err := r.db.Where("sites_id = ?  and id= ? ", idSites, idResource).Delete(entity.Resource{}).Error
	return err
}

func (r *repositoryResource) FindOneResource(idSites int, idResource int) (*entity.Resource, error) {
	var res entity.Resource
	err := r.db.Where("sites_id = ?  and id= ? ", idSites, idResource).First(res).Error
	return &res, err
}
