package usecase_test

import (
	"errors"
	"net/http"
	"service-sites/internal/domain/entity"
	"service-sites/internal/domain/usecase"
	"service-sites/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func Test_SeeSitesOk(t *testing.T) {
	repoSite := mocks.NewIRepositorySites(t)
	sites := []entity.Sites{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	usecaseSites := usecase.NewSitesUseCase(repoSite)
	repoSite.On("FindSites", mock.Anything).Return(&sites, nil)
	_, status := usecaseSites.SeeSites(map[string]interface{}{})
	assert.Equal(t, status, http.StatusOK)
}

func Test_SeeSitesErr(t *testing.T) {
	repoSite := mocks.NewIRepositorySites(t)
	sites := []entity.Sites{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	usecaseSites := usecase.NewSitesUseCase(repoSite)
	repoSite.On("FindSites", mock.Anything).Return(&sites, errors.New(""))
	_, status := usecaseSites.SeeSites(map[string]interface{}{})
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_SeeOneSitesOk(t *testing.T) {
	repoSite := mocks.NewIRepositorySites(t)
	sites := entity.Sites{
		Model: entity.Model{
			ID: 1,
		},
	}

	usecaseSites := usecase.NewSitesUseCase(repoSite)
	repoSite.On("FindOneSites", mock.Anything).Return(&sites, nil)
	_, status := usecaseSites.SeeOneSites(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_SeeOneSitesErr(t *testing.T) {
	repoSite := mocks.NewIRepositorySites(t)
	sites := entity.Sites{
		Model: entity.Model{
			ID: 1,
		},
	}

	usecaseSites := usecase.NewSitesUseCase(repoSite)
	repoSite.On("FindOneSites", mock.Anything).Return(&sites, errors.New(""))
	_, status := usecaseSites.SeeOneSites(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteSitesOk(t *testing.T) {
	repoSite := mocks.NewIRepositorySites(t)
	usecaseSites := usecase.NewSitesUseCase(repoSite)
	repoSite.On("DeleteSites", mock.Anything).Return(nil)
	_, status := usecaseSites.DeleteSites(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteSitesErr(t *testing.T) {
	repoSite := mocks.NewIRepositorySites(t)
	usecaseSites := usecase.NewSitesUseCase(repoSite)
	repoSite.On("DeleteSites", mock.Anything).Return(errors.New(""))
	_, status := usecaseSites.DeleteSites(1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_AddSitesOk(t *testing.T) {
	repoSite := mocks.NewIRepositorySites(t)
	sites := entity.Sites{
		Model: entity.Model{
			ID: 1,
		},
	}

	usecaseSites := usecase.NewSitesUseCase(repoSite)
	repoSite.On("InsertSites", mock.Anything).Return(&sites, nil)
	_, status := usecaseSites.AddSites(sites)
	assert.Equal(t, status, http.StatusOK)
}

func Test_AddSitesErr(t *testing.T) {
	repoSite := mocks.NewIRepositorySites(t)
	sites := entity.Sites{
		Model: entity.Model{
			ID: 1,
		},
	}

	usecaseSites := usecase.NewSitesUseCase(repoSite)
	repoSite.On("InsertSites", mock.Anything).Return(&sites, errors.New(""))
	_, status := usecaseSites.AddSites(sites)
	assert.Equal(t, status, http.StatusBadRequest)
}
