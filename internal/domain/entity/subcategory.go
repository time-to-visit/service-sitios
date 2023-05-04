package entity

import (
	"time"

	"gorm.io/gorm"
)

type Subcategory struct {
	Model
	Nombre      string                   `gorm:"column:nombre;type:varchar(255);not null" json:"nombre" validate:"required"`
	UrlImagen   string                   `gorm:"column:url_imagen;type:text;not null" json:"url_imagen" validate:"required"`
	Descripcion string                   `gorm:"column:desccripcion;type:varchar(255);not null" json:"descripcion" validate:"required"`
	Estado      bool                     `gorm:"column:estado;type:boolean;not null" json:"estado" validate:"required"`
	CategoryId  int                      `gorm:"column:category_id;type:int(11);not null" json:"category_id" validate:"required"`
	Category    *CategoryWithoutValidate `gorm:"joinForeignKey:category_id;foreignKey:id;references:CategoryId" json:"categories,omitempty"`
}

type SubcategoryWihtoutValidate struct {
	Model
	Nombre      string                   `gorm:"column:nombre;type:varchar(255);not null" json:"nombre"`
	UrlImagen   string                   `gorm:"column:url_imagen;type:text;not null" json:"url_imagen"`
	Descripcion string                   `gorm:"column:desccripcion;type:varchar(255);not null" json:"descripcion"`
	Estado      bool                     `gorm:"column:estado;type:boolean;not null" json:"estado" `
	CategoryId  int                      `gorm:"column:category_id;type:int(11);not null" json:"category_id"`
	Category    *CategoryWithoutValidate `gorm:"joinForeignKey:category_id;foreignKey:id;references:CategoryId" json:"categories,omitempty"`
}

func (r *SubcategoryWihtoutValidate) TableName() string {
	return "subcategories"
}

func (m Subcategory) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Subcategory) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
