package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Model
	Nombre      string                       `gorm:"column:nombre;type:varchar(255);not null" json:"nombre" validate:"required"`
	UrlImagen   string                       `gorm:"column:url_imagen;type:varchar(255);not null" json:"url_imagen"`
	Descripcion string                       `gorm:"column:descripcion;type:varchar(255);not null" json:"descripcion" validate:"required"`
	Estado      bool                         `gorm:"column:estado;type:boolean;not null" json:"estado" validate:"required"`
	SubCategory []SubcategoryWihtoutValidate `json:"subcategories"`
}

type CategoryWithoutValidate struct {
	Model
	Nombre      string `gorm:"column:nombre;type:varchar(255);not null" json:"nombre"`
	UrlImagen   string `gorm:"column:url_imagen;type:varchar(255);not null" json:"url_imagen"`
	Descripcion string `gorm:"column:descripcion;type:varchar(255);not null" json:"descripcion"`
	Estado      bool   `gorm:"column:estado;type:boolean;not null" json:"estado"`
}

func (r *CategoryWithoutValidate) TableName() string {
	return "categories"
}

func (m Category) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Category) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
