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

func Test_InsertReviewOk(t *testing.T) {
	cat := entity.Review{
		Model: entity.Model{
			ID: 1,
		},
	}

	site := entity.Sites{
		Model: entity.Model{
			ID: 1,
		},
	}

	repoCat := mocks.NewIRepositoryReview(t)
	repoSite := mocks.NewIRepositorySites(t)

	usecaseCat := usecase.NewReviewUseCase(repoCat, repoSite)
	repoSite.On("FindOneSites", mock.Anything).Return(&site, nil)
	repoCat.On("InsertReview", mock.Anything).Return(&cat, nil)
	_, status := usecaseCat.AddReview(cat)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertReviewErrData(t *testing.T) {
	cat := entity.Review{
		Model: entity.Model{
			ID: 1,
		},
	}

	site := entity.Sites{
		Model: entity.Model{
			ID: 1,
		},
	}

	repoCat := mocks.NewIRepositoryReview(t)
	repoSite := mocks.NewIRepositorySites(t)

	usecaseCat := usecase.NewReviewUseCase(repoCat, repoSite)
	repoSite.On("FindOneSites", mock.Anything).Return(&site, errors.New(""))
	_, status := usecaseCat.AddReview(cat)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertReviewErrInsert(t *testing.T) {
	cat := entity.Review{
		Model: entity.Model{
			ID: 1,
		},
	}

	site := entity.Sites{
		Model: entity.Model{
			ID: 1,
		},
	}

	repoCat := mocks.NewIRepositoryReview(t)
	repoSite := mocks.NewIRepositorySites(t)

	usecaseCat := usecase.NewReviewUseCase(repoCat, repoSite)
	repoSite.On("FindOneSites", mock.Anything).Return(&site, nil)
	repoCat.On("InsertReview", mock.Anything).Return(&cat, errors.New(""))
	_, status := usecaseCat.AddReview(cat)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteReviewOk(t *testing.T) {

	repoCat := mocks.NewIRepositoryReview(t)
	repoSite := mocks.NewIRepositorySites(t)

	usecaseCat := usecase.NewReviewUseCase(repoCat, repoSite)
	repoCat.On("DeleteReview", mock.Anything, mock.Anything).Return(nil)
	_, status := usecaseCat.DeleteReview(1, 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteReviewErr(t *testing.T) {

	repoCat := mocks.NewIRepositoryReview(t)
	repoSite := mocks.NewIRepositorySites(t)

	usecaseCat := usecase.NewReviewUseCase(repoCat, repoSite)
	repoCat.On("DeleteReview", mock.Anything, mock.Anything).Return(errors.New(""))
	_, status := usecaseCat.DeleteReview(1, 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_SeeReviewOk(t *testing.T) {

	repoCat := mocks.NewIRepositoryReview(t)
	repoSite := mocks.NewIRepositorySites(t)
	cat := []entity.Review{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	usecaseCat := usecase.NewReviewUseCase(repoCat, repoSite)
	repoCat.On("FindReview", mock.Anything).Return(&cat, nil)
	_, status := usecaseCat.SeeReview(1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_SeeReviewErr(t *testing.T) {

	repoCat := mocks.NewIRepositoryReview(t)
	repoSite := mocks.NewIRepositorySites(t)
	cat := []entity.Review{
		{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
	usecaseCat := usecase.NewReviewUseCase(repoCat, repoSite)
	repoCat.On("FindReview", mock.Anything).Return(&cat, errors.New(""))
	_, status := usecaseCat.SeeReview(1)
	assert.Equal(t, status, http.StatusBadRequest)
}
