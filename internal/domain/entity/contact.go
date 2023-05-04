package entity

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	Model
	IsWp    bool                  `gorm:"column:is_wp;type:boolean;not null" json:"is_wp" validate:"required"`
	Nombre  string                `gorm:"column:nombre;type:varchar(255);not null" json:"nombre" validate:"required"`
	Numero  string                `gorm:"column:numero;type:varchar(255);not null" json:"numero" validate:"required"`
	SitesId int                   `gorm:"column:sites_id;type:int(11);not null" json:"sites_id" validate:"required"`
	Sites   *SitesWithoutValidate `gorm:"joinForeignKey:sites_id;foreignKey:id;references:SitesId" json:"sites,omitempty"`
	Estado  bool                  `gorm:"column:estado;type:boolean;not null" json:"estado" validate:"required"`
}

type ContactWithoutValidate struct {
	Model
	IsWp    bool                  `gorm:"column:is_wp;type:boolean;not null" json:"is_wp"`
	Nombre  string                `gorm:"column:nombre;type:varchar(255);not null" json:"nombre"`
	Numero  string                `gorm:"column:numero;type:varchar(255);not null" json:"numero"`
	SitesId int                   `gorm:"column:sites_id;type:int(11);not null" json:"sites_id"`
	Sites   *SitesWithoutValidate `gorm:"joinForeignKey:sites_id;foreignKey:id;references:SitesId" json:"sites,omitempty"`
	Estado  bool                  `gorm:"column:estado;type:boolean;not null" json:"estado"`
}

func (r *ContactWithoutValidate) TableName() string {
	return "contacts"
}

func (m Contact) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Contact) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
