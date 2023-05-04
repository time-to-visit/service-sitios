package entity

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	Model
	Nombre         string                          `gorm:"column:nombre;type:varchar(255);not null" json:"nombre" validate:"required"`
	Bandera        string                          `gorm:"column:bandera;type:varchar(255);not null" json:"bandera" validate:"required"`
	Estado         bool                            `gorm:"column:estado;type:boolean;not null" json:"estado" validate:"required"`
	Municipalities []MunicipalitiesWithoutValidate ` json:"municipios"`
}

type DepartmentWithoutValidate struct {
	Model
	Nombre  string `gorm:"column:nombre;type:varchar(255);not null" json:"nombre"`
	Bandera string `gorm:"column:bandera;type:varchar(255);not null" json:"bandera"`
	Estado  bool   `gorm:"column:estado;type:boolean;not null" json:"estado"`
}

func (m DepartmentWithoutValidate) TableName() string {
	return "departments"
}

func (m Department) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Department) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
