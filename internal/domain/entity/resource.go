package entity

import (
	"time"

	"gorm.io/gorm"
)

type Resource struct {
	Model
	TypeRecursos string               `gorm:"column:type_recursos;type:varchar(255);not null" json:"type_resource" validate:"required"`
	Payload      string               `gorm:"column:payload;type:varchar(255);not null" json:"payload" validate:"required"`
	Estado       bool                 `gorm:"column:estado;type:boolean;not null" json:"estado" validate:"required"`
	SitesId      int                  `gorm:"column:sites_id;type:int(11);not null" json:"sites_id" validate:"required"`
	Sites        SitesWithoutValidate `gorm:"joinForeignKey:sites_id;foreignKey:id;references:SitesId"`
}

type ResourceWithoutValidate struct {
	Model
	TypeRecursos string `gorm:"column:type_recursos;type:varchar(255);not null" json:"type_resource"`
	Payload      string `gorm:"column:payload;type:varchar(255);not null" json:"payload"`
	Estado       bool   `gorm:"column:estado;type:boolean;not null" json:"estado" `
	SitesId      int    `gorm:"column:sites_id;type:int(11);not null" json:"sites_id"`
}

func (res ResourceWithoutValidate) TableName() string {
	return "resources"
}

func (m Resource) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Resource) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
