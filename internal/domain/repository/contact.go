package repository

import (
	"service-sites/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryContact interface {
	InsertContact(contact entity.Contact) (*entity.Contact, error)
	DeleteContact(idSites int, idContact int) error
}

type repositoryContact struct {
	db *gorm.DB
}

func NewRepositoryContact(db *gorm.DB) IRepositoryContact {
	return &repositoryContact{
		db,
	}
}

func (r *repositoryContact) InsertContact(contact entity.Contact) (*entity.Contact, error) {
	err := r.db.Create(&contact).Error
	return &contact, err
}

func (r *repositoryContact) DeleteContact(idSites int, idContact int) error {
	err := r.db.Where("sites_id = ?  and id= ? ", idSites, idContact).Delete(entity.Contact{}).Error
	return err
}
