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

func Test_InsertResourceOk(t *testing.T) {
	cat := entity.Resource{
		TypeRecursos: "IMAGE",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoCat := mocks.NewIRepositoryResource(t)

	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourceUseCase(repoCat, objectFile)
	repoCat.On("InsertResource", mock.Anything).Return(&cat, nil)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.AddResource(context.Background(), cat)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertResourceErrorFile(t *testing.T) {
	cat := entity.Resource{
		TypeRecursos: "IMAGE",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoCat := mocks.NewIRepositoryResource(t)

	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourceUseCase(repoCat, objectFile)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	_, status := usecaseCat.AddResource(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertResourceErrData(t *testing.T) {
	cat := entity.Resource{
		TypeRecursos: "IMAGE",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoCat := mocks.NewIRepositoryResource(t)

	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourceUseCase(repoCat, objectFile)
	repoCat.On("InsertResource", mock.Anything).Return(&cat, errors.New(""))
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.AddResource(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteResourceOk(t *testing.T) {
	cat := entity.Resource{
		TypeRecursos: "IMAGE",
		Payload:      "https://asdasdasd.com/asdasd",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoCat := mocks.NewIRepositoryResource(t)

	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourceUseCase(repoCat, objectFile)
	repoCat.On("FindOneResource", mock.Anything, mock.Anything).Return(&cat, nil)
	repoCat.On("DeleteResource", mock.Anything, mock.Anything).Return(nil)

	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.DeleteResource(context.Background(), 1, 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteResourceErrFile(t *testing.T) {
	cat := entity.Resource{
		TypeRecursos: "IMAGE",
		Payload:      "https://asdasdasd.com/asdasd",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoCat := mocks.NewIRepositoryResource(t)

	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourceUseCase(repoCat, objectFile)
	repoCat.On("FindOneResource", mock.Anything, mock.Anything).Return(&cat, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))
	_, status := usecaseCat.DeleteResource(context.Background(), 1, 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteResourceErrData(t *testing.T) {
	cat := entity.Resource{
		TypeRecursos: "IMAGE",
		Payload:      "https://asdasdasd.com/asdasd",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoCat := mocks.NewIRepositoryResource(t)

	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourceUseCase(repoCat, objectFile)
	repoCat.On("FindOneResource", mock.Anything, mock.Anything).Return(&cat, errors.New(""))
	_, status := usecaseCat.DeleteResource(context.Background(), 1, 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteResourceErrDel(t *testing.T) {
	cat := entity.Resource{
		TypeRecursos: "IMAGE",
		Payload:      "https://asdasdasd.com/asdasd",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoCat := mocks.NewIRepositoryResource(t)

	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewResourceUseCase(repoCat, objectFile)
	repoCat.On("FindOneResource", mock.Anything, mock.Anything).Return(&cat, nil)
	repoCat.On("DeleteResource", mock.Anything, mock.Anything).Return(errors.New(""))

	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	_, status := usecaseCat.DeleteResource(context.Background(), 1, 1)
	assert.Equal(t, status, http.StatusBadRequest)
}
