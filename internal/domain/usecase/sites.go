package usecase

import (
	"net/http"
	"service-sites/internal/domain/entity"
	objectValues "service-sites/internal/domain/object_values"
	"service-sites/internal/domain/repository"
)

type SitesUseCase struct {
	repoSite repository.IRepositorySites
}

func NewSitesUseCase(repoSite repository.IRepositorySites) SitesUseCase {
	return SitesUseCase{
		repoSite,
	}
}

func (e *SitesUseCase) SeeSites(filter map[string]interface{}) (interface{}, int) {
	sites, err := e.repoSite.FindSites(filter)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", sites), http.StatusOK
}

func (e *SitesUseCase) SeeOneSites(idSites int) (interface{}, int) {
	site, err := e.repoSite.FindOneSites(idSites)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", site), http.StatusOK
}

func (e *SitesUseCase) DeleteSites(idSites int) (interface{}, int) {
	err := e.repoSite.DeleteSites(idSites)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el sitio", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sitio eliminado exitosamente", nil), http.StatusOK
}

func (e *SitesUseCase) AddSites(site entity.Sites) (interface{}, int) {
	newSite, err := e.repoSite.InsertSites(site)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar el sitio", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sitio insertada exitosamente", newSite), http.StatusOK
}
