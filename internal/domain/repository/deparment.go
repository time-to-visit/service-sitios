package repository

import (
	"service-sites/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryDepartment interface {
	InsertDepartment(department entity.Department) (*entity.Department, error)
	FindDepartment() (*[]entity.Department, error)
	FindDepartmentById(Id int) (*entity.Department, error)
	DeleteDepartment(idDeparment int) error
}

type repositoryDepartment struct {
	db *gorm.DB
}

func NewRepositoryDepartment(db *gorm.DB) IRepositoryDepartment {
	return &repositoryDepartment{
		db,
	}
}

func (r *repositoryDepartment) InsertDepartment(department entity.Department) (*entity.Department, error) {
	err := r.db.Create(&department).Error
	return &department, err
}

func (r *repositoryDepartment) FindDepartment() (*[]entity.Department, error) {
	var entitys []entity.Department
	err := r.db.Preload("Municipalities").Find(&entitys).Error
	return &entitys, err
}

func (r *repositoryDepartment) FindDepartmentById(Id int) (*entity.Department, error) {
	var deparment entity.Department
	err := r.db.First(&deparment, Id).Error
	return &deparment, err
}

func (r *repositoryDepartment) DeleteDepartment(idDeparment int) error {
	return r.db.Delete(entity.Department{}, idDeparment).Error
}
