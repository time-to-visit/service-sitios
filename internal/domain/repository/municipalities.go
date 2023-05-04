package repository

import (
	"service-sites/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryMunicipalities interface {
	InsertMunicipalitiest(municipalities entity.Municipalities) (*entity.Municipalities, error)
	FindMunicipalities(idDeparment int) (*[]entity.Municipalities, error)
	DeleteMunicipalities(idMunicipalities int) error
	FindOneMunicipalities(idMunicipalities int) (*entity.Municipalities, error)
}

type repositoryMunicipalities struct {
	db *gorm.DB
}

func NewRepositoryMunicipalities(db *gorm.DB) IRepositoryMunicipalities {
	return &repositoryMunicipalities{
		db,
	}
}

func (r *repositoryMunicipalities) InsertMunicipalitiest(municipalities entity.Municipalities) (*entity.Municipalities, error) {
	err := r.db.Create(&municipalities).Error
	return &municipalities, err
}

func (r *repositoryMunicipalities) FindMunicipalities(idDeparment int) (*[]entity.Municipalities, error) {
	var municipalities []entity.Municipalities
	err := r.db.Where("department_id = ?", idDeparment).Find(&municipalities).Error
	return &municipalities, err
}

func (r *repositoryMunicipalities) DeleteMunicipalities(idMunicipalities int) error {
	return r.db.Delete(entity.Municipalities{}, idMunicipalities).Error
}

func (r *repositoryMunicipalities) FindOneMunicipalities(idMunicipalities int) (*entity.Municipalities, error) {
	var municipality entity.Municipalities
	err := r.db.First(&municipality, idMunicipalities).Error
	return &municipality, err
}
