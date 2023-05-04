package repository

import (
	"service-sites/internal/domain/entity"

	"gorm.io/gorm"
)

type IRepositoryReview interface {
	InsertReview(review entity.Review) (*entity.Review, error)
	FindReview(idSites int) (*[]entity.Review, error)
	DeleteReview(idSites int, idReview int) error
}

type repositoryReview struct {
	db *gorm.DB
}

func NewRepositoryReview(db *gorm.DB) IRepositoryReview {
	return &repositoryReview{
		db,
	}
}

func (r *repositoryReview) InsertReview(review entity.Review) (*entity.Review, error) {
	err := r.db.Create(&review).Error
	return &review, err
}
func (r *repositoryReview) FindReview(idSites int) (*[]entity.Review, error) {
	var reviews []entity.Review
	err := r.db.Where("sites_id = ?", idSites).Find(&reviews).Error
	return &reviews, err
}
func (r *repositoryReview) DeleteReview(idSites int, idReview int) error {
	err := r.db.Where("sites_id = ?  and id= ? ", idSites, idReview).Delete(entity.Review{}).Error
	return err
}
