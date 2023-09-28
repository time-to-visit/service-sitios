package entity

import (
	"time"

	"gorm.io/gorm"
)

type Sites struct {
	Model
	CategoryID       int                            `gorm:"column:category_id;type:int(11);not null" json:"category_id" validate:"required"`
	SubCategoryID    int                            `gorm:"column:subcategory_id;type:int(11);not null" json:"subcategory_id" validate:"required"`
	SubCategory      *SubcategoryWihtoutValidate    `gorm:"joinForeignKey:subcategory_id;foreignKey:id;references:SubCategoryID" json:"subcategory,omitempty"`
	Category         *CategoryWithoutValidate       `gorm:"joinForeignKey:category_id;foreignKey:id;references:CategoryID" json:"category,omitempty"`
	MunicipalitiesID int                            `gorm:"column:municipalities_id;type:int(11);not null" json:"municipalities_id" validate:"required"`
	Municipalities   *MunicipalitiesWithoutValidate `gorm:"joinForeignKey:municipalities_id;foreignKey:id;references:MunicipalitiesID" json:"municipalities,omitempty"`
	Nombre           string                         `gorm:"column:nombre;type:varchar(255);not null" json:"nombre" validate:"required"`
	Descripcion      string                         `gorm:"column:descripcion;type:text;size:65535;not null" json:"descripcion" validate:"required"`
	Dirrecion        string                         `gorm:"column:dirrecion;type:varchar(255);not null" json:"dirrecion" validate:"required"`
	Puntuacion       string                         `gorm:"column:puntuacion;type:varchar(255);not null" json:"puntuacion" validate:"required"`
	Estado           bool                           `gorm:"column:estado;type:boolean;not null" json:"estado" validate:"required"`
	Latitud          string                         `gorm:"column:latitud;type:varchar(255);not null" json:"latitud" validate:"required"`
	Longitud         string                         `gorm:"column:longitud;type:varchar(255);not null" json:"longitud" validate:"required"`
	NumeroResena     string                         `gorm:"column:numero_resena;type:varchar(255);not null" json:"numero_resena" validate:"required"`
	Contact          []ContactWithoutValidate       `json:"contacts"`
	Resource         []ResourceWithoutValidate      `json:"resource"`
}

type SitesWithoutValidate struct {
	Model
	CategoryID       int                            `gorm:"column:category_id;type:int(11);not null" json:"category_id"`
	Category         *CategoryWithoutValidate       `gorm:"joinForeignKey:category_id;foreignKey:id;references:CategoryID"`
	SubCategoryID    int                            `gorm:"column:subcategory_id;type:int(11);not null" json:"subcategory_id"`
	MunicipalitiesID int                            `gorm:"column:municipalities_id;type:int(11);not null" json:"municipalities_id"`
	Municipalities   *MunicipalitiesWithoutValidate `gorm:"joinForeignKey:municipalities_id;foreignKey:id;references:MunicipalitiesID" json:"municipalities,omitempty"`
	Nombre           string                         `gorm:"column:nombre;type:varchar(255);not null" json:"nombre"`
	Descripcion      string                         `gorm:"column:descripcion;type:varchar(255);not null" json:"descripcion"`
	Dirrecion        string                         `gorm:"column:dirrecion;type:varchar(255);not null" json:"dirrecion"`
	Puntuacion       string                         `gorm:"column:puntuacion;type:varchar(255);not null" json:"puntuacion"`
	Estado           bool                           `gorm:"column:estado;type:boolean;not null" json:"estado"`
	Latitud          string                         `gorm:"column:latitud;type:varchar(255);not null" json:"latitud"`
	Longitud         string                         `gorm:"column:longitud;type:varchar(255);not null" json:"longitud"`
	NumeroResena     string                         `gorm:"column:numero_resena;type:varchar(255);not null" json:"numero_resena"`
}

func (m SitesWithoutValidate) TableName() string {
	return "sites"
}

func (m Sites) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m Sites) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now()
	return
}
