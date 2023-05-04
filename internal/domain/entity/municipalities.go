package entity

import (
	"time"

	"gorm.io/gorm"
)

type Municipalities struct {
	Model
	Nombre       string                     `gorm:"column:nombre;type:varchar(255);not null" json:"nombre" validate:"required"`
	Bandera      string                     `gorm:"column:bandera;type:varchar(255);not null" json:"bandera" validate:"required"`
	Estado       bool                       `gorm:"column:estado;type:boolean;not null" json:"estado" validate:"required"`
	DepartmentId int                        `gorm:"column:department_id;type:int(11);not null" json:"department_id" validate:"required"`
	Department   *DepartmentWithoutValidate `gorm:"joinForeignKey:department_id;foreignKey:id;references:DepartmentId" json:"department,omitempty"`
}

type MunicipalitiesWithoutValidate struct {
	Model
	Nombre       string `gorm:"column:nombre;type:varchar(255);not null" json:"nombre" `
	Bandera      string `gorm:"column:bandera;type:varchar(255);not null" json:"bandera" `
	Estado       bool   `gorm:"column:estado;type:boolean;not null" json:"estado" `
	DepartmentId int    `gorm:"column:department_id;type:int(11);not null" json:"_" `
}

func (r *MunicipalitiesWithoutValidate) TableName() string {
	return "municipalities"
}

func (m Municipalities) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Municipalities) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
