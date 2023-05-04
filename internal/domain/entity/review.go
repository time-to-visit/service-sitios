package entity

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	Model
	UserId        int    `gorm:"column:user_id;type:int(11);not null"  json:"user_id"`
	UserNombre    string `gorm:"column:user_nombre;type:varchar(255);not null" json:"user_nombre"`
	UrlImagenUser string `gorm:"column:url_imagen_user;type:varchar(255);not null" json:"url_imagen_user" `
	Comment       string `gorm:"column:comment;type:varchar(255);not null" json:"comment" validate:"required"`
	Valoracion    string `gorm:"column:valoracion;type:varchar(255);not null" json:"valoracion" validate:"required"`
	SitesID       int    `gorm:"column:sites_id;type:varchar(255);not null" json:"sites_id" validate:"required"`
}

func (m Review) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Review) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
