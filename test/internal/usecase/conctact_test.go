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

func Test_InsertContactOk(t *testing.T) {
	repoCat := mocks.NewIRepositoryContact(t)
	cats := entity.Contact{
		Nombre: "dasdasd",
	}
	usecaseCat := usecase.NewContactUseCase(repoCat)

	repoCat.On("InsertContact", mock.Anything).Return(&cats, nil)

	_, status := usecaseCat.InsertContact(cats)
	assert.Equal(t, http.StatusOK, status)
}

func Test_InsertContactErr(t *testing.T) {
	repoCat := mocks.NewIRepositoryContact(t)
	cats := entity.Contact{
		Nombre: "dasdasd",
	}
	usecaseCat := usecase.NewContactUseCase(repoCat)

	repoCat.On("InsertContact", mock.Anything).Return(&cats, errors.New(""))

	_, status := usecaseCat.InsertContact(cats)
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_DeleteContactOk(t *testing.T) {
	repoCat := mocks.NewIRepositoryContact(t)
	usecaseCat := usecase.NewContactUseCase(repoCat)

	repoCat.On("DeleteContact", mock.Anything, mock.Anything).Return(nil)

	_, status := usecaseCat.DeleteContact(1, 1)
	assert.Equal(t, http.StatusOK, status)
}

func Test_DeleteContactErr(t *testing.T) {
	repoCat := mocks.NewIRepositoryContact(t)
	usecaseCat := usecase.NewContactUseCase(repoCat)

	repoCat.On("DeleteContact", mock.Anything, mock.Anything).Return(errors.New(""))

	_, status := usecaseCat.DeleteContact(1, 1)
	assert.Equal(t, http.StatusBadRequest, status)
}
