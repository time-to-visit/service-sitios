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

func Test_InsertSubcategory(t *testing.T) {
	sub := entity.Subcategory{
		Model: entity.Model{
			ID: 1,
		},
	}
	cat := entity.Category{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoSub := mocks.NewIRepositorySubCategory(t)
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoCat.On("FindCategoryById", mock.Anything).Return(&cat, nil)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoSub.On("InsertSubcategory", mock.Anything).Return(&sub, nil)
	usecaseSub := usecase.NewSubCategoryUseCase(repoSub, repoCat, objectFile)
	_, status := usecaseSub.InsertSubcategory(context.Background(), sub)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertSubcategoryErrSetFile(t *testing.T) {
	sub := entity.Subcategory{
		Model: entity.Model{
			ID: 1,
		},
	}
	cat := entity.Category{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoSub := mocks.NewIRepositorySubCategory(t)
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoCat.On("FindCategoryById", mock.Anything).Return(&cat, nil)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	usecaseSub := usecase.NewSubCategoryUseCase(repoSub, repoCat, objectFile)
	_, status := usecaseSub.InsertSubcategory(context.Background(), sub)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertSubcategoryErrFindCategory(t *testing.T) {
	sub := entity.Subcategory{
		Model: entity.Model{
			ID: 1,
		},
	}
	cat := entity.Category{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoSub := mocks.NewIRepositorySubCategory(t)
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoCat.On("FindCategoryById", mock.Anything).Return(&cat, errors.New(""))
	usecaseSub := usecase.NewSubCategoryUseCase(repoSub, repoCat, objectFile)
	_, status := usecaseSub.InsertSubcategory(context.Background(), sub)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertSubcategoryErrInsert(t *testing.T) {
	sub := entity.Subcategory{
		Model: entity.Model{
			ID: 1,
		},
	}
	cat := entity.Category{
		Model: entity.Model{
			ID: 1,
		},
	}
	repoSub := mocks.NewIRepositorySubCategory(t)
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoCat.On("FindCategoryById", mock.Anything).Return(&cat, nil)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	repoSub.On("InsertSubcategory", mock.Anything).Return(&sub, errors.New(""))
	usecaseSub := usecase.NewSubCategoryUseCase(repoSub, repoCat, objectFile)
	_, status := usecaseSub.InsertSubcategory(context.Background(), sub)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteSubcategory(t *testing.T) {
	sub := entity.Subcategory{
		UrlImagen: "http://asdasda.com/asda",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoSub := mocks.NewIRepositorySubCategory(t)
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoSub.On("FindOneSubcategory", mock.Anything).Return(&sub, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoSub.On("DeleteSubcategory", mock.Anything, mock.Anything).Return(nil)
	usecaseSub := usecase.NewSubCategoryUseCase(repoSub, repoCat, objectFile)
	_, status := usecaseSub.DeleteSubCategory(context.Background(), 1, 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteSubcategoryErrFindOne(t *testing.T) {
	sub := entity.Subcategory{
		UrlImagen: "http://asdasda.com/asda",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoSub := mocks.NewIRepositorySubCategory(t)
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoSub.On("FindOneSubcategory", mock.Anything).Return(&sub, errors.New(""))
	usecaseSub := usecase.NewSubCategoryUseCase(repoSub, repoCat, objectFile)
	_, status := usecaseSub.DeleteSubCategory(context.Background(), 1, 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteSubcategoryErrDeleteFile(t *testing.T) {
	sub := entity.Subcategory{
		UrlImagen: "http://asdasda.com/asda",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoSub := mocks.NewIRepositorySubCategory(t)
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoSub.On("FindOneSubcategory", mock.Anything).Return(&sub, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))
	usecaseSub := usecase.NewSubCategoryUseCase(repoSub, repoCat, objectFile)
	_, status := usecaseSub.DeleteSubCategory(context.Background(), 1, 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteSubcategoryErrDelete(t *testing.T) {
	sub := entity.Subcategory{
		UrlImagen: "http://asdasda.com/asda",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoSub := mocks.NewIRepositorySubCategory(t)
	repoCat := mocks.NewIRepositoryCategory(t)
	objectFile := mocks.NewIGCImageRepo(t)
	repoSub.On("FindOneSubcategory", mock.Anything).Return(&sub, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoSub.On("DeleteSubcategory", mock.Anything, mock.Anything).Return(errors.New(""))
	usecaseSub := usecase.NewSubCategoryUseCase(repoSub, repoCat, objectFile)
	_, status := usecaseSub.DeleteSubCategory(context.Background(), 1, 1)
	assert.Equal(t, status, http.StatusBadRequest)
}
