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

func Test_SeeDepartmentOk(t *testing.T) {
	repoCat := mocks.NewIRepositoryDepartment(t)
	cats := []entity.Department{
		{
			Nombre: "dasdasd",
		},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewDepartmentUseCase(repoCat, objectFile)

	repoCat.On("FindDepartment").Return(&cats, nil)

	_, status := usecaseCat.SeeDeparment()
	assert.Equal(t, http.StatusOK, status)
}

func Test_SeeDepartmentErr(t *testing.T) {
	repoCat := mocks.NewIRepositoryDepartment(t)
	cats := []entity.Department{
		{
			Nombre: "dasdasd",
		},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewDepartmentUseCase(repoCat, objectFile)

	repoCat.On("FindDepartment").Return(&cats, errors.New("err"))

	_, status := usecaseCat.SeeDeparment()
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_DeleteDepartmentOk(t *testing.T) {
	cat := entity.Department{
		Nombre:  "dasdasd",
		Bandera: "http://dadas.co,/dadasdadasdsa",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryDepartment(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewDepartmentUseCase(repoCat, objectFile)
	repoCat.On("FindDepartmentById", mock.Anything).Return(&cat, nil)
	repoCat.On("DeleteDepartment", mock.Anything).Return(nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	_, status := usecaseCat.DeleteDeparment(context.Background(), 1)
	assert.Equal(t, status, http.StatusOK)
}

func Test_DeleteDepartmentFindErr(t *testing.T) {
	cat := entity.Department{
		Nombre:  "dasdasd",
		Bandera: "http://dadas.co,/dadasdadasdsa",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryDepartment(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewDepartmentUseCase(repoCat, objectFile)
	repoCat.On("FindDepartmentById", mock.Anything).Return(&cat, errors.New(""))
	_, status := usecaseCat.DeleteDeparment(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteDepartmentObjectErr(t *testing.T) {
	cat := entity.Department{
		Nombre:  "dasdasd",
		Bandera: "http://dadas.co,/dadasdadasdsa",

		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryDepartment(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewDepartmentUseCase(repoCat, objectFile)
	repoCat.On("FindDepartmentById", mock.Anything).Return(&cat, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))
	_, status := usecaseCat.DeleteDeparment(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)

}

func Test_DeleteDepartmentDeleteErr(t *testing.T) {
	cat := entity.Department{
		Nombre:  "dasdasd",
		Bandera: "http://dadas.co,/dadasdadasdsa",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryDepartment(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewDepartmentUseCase(repoCat, objectFile)
	repoCat.On("FindDepartmentById", mock.Anything).Return(&cat, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repoCat.On("DeleteDepartment", mock.Anything).Return(errors.New(""))
	_, status := usecaseCat.DeleteDeparment(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertDepartmentOk(t *testing.T) {
	cat := entity.Department{
		Nombre: "dasdasd",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryDepartment(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewDepartmentUseCase(repoCat, objectFile)
	repoCat.On("InsertDepartment", mock.Anything).Return(&cat, nil)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.InsertDeparment(context.Background(), cat)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertDepartmentObjectErr(t *testing.T) {
	cat := entity.Department{
		Nombre: "dasdasd",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryDepartment(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewDepartmentUseCase(repoCat, objectFile)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	_, status := usecaseCat.InsertDeparment(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertDepartmentInsertErr(t *testing.T) {
	cat := entity.Department{
		Nombre: "dasdasd",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryDepartment(t)
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewDepartmentUseCase(repoCat, objectFile)
	repoCat.On("InsertDepartment", mock.Anything).Return(&cat, errors.New(""))
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.InsertDeparment(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}
