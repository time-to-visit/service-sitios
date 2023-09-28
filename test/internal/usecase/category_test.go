package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"service-sites/internal/domain/entity"
	"service-sites/internal/domain/usecase"
	"service-sites/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func Test_SeeCategoryOk(t *testing.T) {
	repoCat := mocks.NewIRepositoryCategory(t)
	cats := []entity.Category{
		{
			Nombre: "dasdasd",
		},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewCategoryUseCase(repoCat, objectFile)

	repoCat.On("FindCategory").Return(&cats, nil)

	_, status := usecaseCat.SeeCategory()
	assert.Equal(t, http.StatusOK, status)
}

func Test_SeeCategoryErr(t *testing.T) {
	repoCat := mocks.NewIRepositoryCategory(t)
	cats := []entity.Category{
		{
			Nombre: "dasdasd",
		},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewCategoryUseCase(repoCat, objectFile)

	repoCat.On("FindCategory").Return(&cats, errors.New("err"))

	_, status := usecaseCat.SeeCategory()
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_UpdateCategoryOk(t *testing.T) {
	cat := entity.Category{
		Nombre: "dasdasd",
	}
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewCategoryUseCase(repoCat, objectFile)
	repoCat.On("UpdateCategory", mock.Anything).Return(&cat, nil)

	_, status := usecaseCat.UpdateCategory(cat)
	assert.Equal(t, http.StatusOK, status)
}

func Test_UpdateCategoryErr(t *testing.T) {
	cat := entity.Category{
		Nombre: "dasdasd",
	}
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewCategoryUseCase(repoCat, objectFile)
	repoCat.On("UpdateCategory", mock.Anything).Return(&cat, errors.New(""))

	_, status := usecaseCat.UpdateCategory(cat)
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_DeleteCategoryOk(t *testing.T) {
	cat := entity.Category{
		Nombre:    "dasdasd",
		UrlImagen: "https://google.com/dasdasdasdas",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewCategoryUseCase(repoCat, objectFile)
	repoCat.On("FindCategoryById", mock.Anything).Return(&cat, nil)
	repoCat.On("DeleteCategory", mock.Anything).Return(nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	_, status := usecaseCat.DeleteCategory(context.Background(), 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteCategoryFindErr(t *testing.T) {
	cat := entity.Category{
		Nombre:    "dasdasd",
		UrlImagen: "https://google.com/dasdasdasdas",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewCategoryUseCase(repoCat, objectFile)
	repoCat.On("FindCategoryById", mock.Anything).Return(&cat, errors.New(""))
	_, status := usecaseCat.DeleteCategory(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteCategoryObjectErr(t *testing.T) {
	cat := entity.Category{
		Nombre:    "dasdasd",
		UrlImagen: "https://google.com/dasdasdasdas",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewCategoryUseCase(repoCat, objectFile)
	repoCat.On("FindCategoryById", mock.Anything).Return(&cat, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))
	_, status := usecaseCat.DeleteCategory(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)

}

func Test_DeleteCategoryDeleteErr(t *testing.T) {
	cat := entity.Category{
		Nombre:    "dasdasd",
		UrlImagen: "https://google.com/dasdasdasdas",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewCategoryUseCase(repoCat, objectFile)
	repoCat.On("FindCategoryById", mock.Anything).Return(&cat, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoCat.On("DeleteCategory", mock.Anything).Return(errors.New(""))
	_, status := usecaseCat.DeleteCategory(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertCategoryOk(t *testing.T) {
	cat := entity.Category{
		Nombre:    "dasdasd",
		UrlImagen: "https://google.com/dasdasdasdas",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewCategoryUseCase(repoCat, objectFile)
	repoCat.On("RegisterCategory", mock.Anything).Return(&cat, nil)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.InsertCategory(context.Background(), cat)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertCategoryObjectErr(t *testing.T) {
	cat := entity.Category{
		Nombre:    "dasdasd",
		UrlImagen: "https://google.com/dasdasdasdas",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewCategoryUseCase(repoCat, objectFile)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	_, status := usecaseCat.InsertCategory(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertCategoryInsertErr(t *testing.T) {
	cat := entity.Category{
		Nombre:    "dasdasd",
		UrlImagen: "https://google.com/dasdasdasdas",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewCategoryUseCase(repoCat, objectFile)
	repoCat.On("RegisterCategory", mock.Anything).Return(&cat, errors.New(""))
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.InsertCategory(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}
