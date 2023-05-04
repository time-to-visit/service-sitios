package usecase

import (
	"net/http"
	"service-sites/internal/domain/entity"
	objectValues "service-sites/internal/domain/object_values"
	"service-sites/internal/domain/repository"
)

type ReviewUseCase struct {
	repoRev  repository.IRepositoryReview
	repoSite repository.IRepositorySites
}

func NewReviewUseCase(repoRev repository.IRepositoryReview, repoSite repository.IRepositorySites) ReviewUseCase {
	return ReviewUseCase{
		repoRev,
		repoSite,
	}
}

func (e *ReviewUseCase) AddReview(review entity.Review) (interface{}, int) {
	sites, err := e.repoSite.FindOneSites(review.SitesID)
	if err != nil || sites.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "sites no encontrado", nil), http.StatusBadRequest
	}
	newReview, err := e.repoRev.InsertReview(review)
	if err != nil || sites.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al insertar review", nil), http.StatusBadRequest
	}

	return objectValues.NewResponseWithData(http.StatusOK, "ok", "reseña insertada exitosamente", newReview), http.StatusOK
}

func (e *ReviewUseCase) DeleteReview(idSites int, idReview int) (interface{}, int) {
	err := e.repoRev.DeleteReview(idSites, idReview)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un error al eliminar la reseña", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "reseña eliminado exitosamente", nil), http.StatusOK
}

func (e *ReviewUseCase) SeeReview(idSites int) (interface{}, int) {
	review, err := e.repoRev.FindReview(idSites)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un error al traer la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", review), http.StatusOK
}
