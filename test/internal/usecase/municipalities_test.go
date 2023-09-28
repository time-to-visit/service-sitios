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

func Test_SeeMunicipalitiesOk(t *testing.T) {
	repoCat := mocks.NewIRepositoryMunicipalities(t)
	repoDep := mocks.NewIRepositoryDepartment(t)
	cats := []entity.Municipalities{
		{
			Nombre: "dasdasd",
		},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewMunicipalitiesUseCase(repoCat, repoDep, objectFile)

	repoCat.On("FindMunicipalities", mock.Anything).Return(&cats, nil)

	_, status := usecaseCat.SeeMuncipalities(1)
	assert.Equal(t, http.StatusOK, status)
}

func Test_SeeMunicipalitiesErr(t *testing.T) {
	repoCat := mocks.NewIRepositoryMunicipalities(t)
	repoDep := mocks.NewIRepositoryDepartment(t)
	cats := []entity.Municipalities{
		{
			Nombre: "dasdasd",
		},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewMunicipalitiesUseCase(repoCat, repoDep, objectFile)

	repoCat.On("FindMunicipalities", mock.Anything).Return(&cats, errors.New(""))

	_, status := usecaseCat.SeeMuncipalities(1)
	assert.Equal(t, http.StatusBadRequest, status)
}

func Test_InsertMunicipalitiesOk(t *testing.T) {
	cat := entity.Municipalities{
		Nombre: "dasdasd",
		Model: entity.Model{
			ID: 1,
		},
	}

	dep := entity.Department{
		Nombre: "dasdasd",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryMunicipalities(t)
	repoDep := mocks.NewIRepositoryDepartment(t)

	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewMunicipalitiesUseCase(repoCat, repoDep, objectFile)
	repoCat.On("InsertMunicipalitiest", mock.Anything).Return(&cat, nil)
	repoDep.On("FindDepartmentById", mock.Anything).Return(&dep, nil)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.InsertMuncipalities(context.Background(), cat)
	assert.Equal(t, status, http.StatusOK)
}

func Test_InsertMunicipalitiesErrSetFile(t *testing.T) {
	cat := entity.Municipalities{
		Nombre: "dasdasd",
		Model: entity.Model{
			ID: 1,
		},
	}

	repoCat := mocks.NewIRepositoryMunicipalities(t)
	repoDep := mocks.NewIRepositoryDepartment(t)

	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewMunicipalitiesUseCase(repoCat, repoDep, objectFile)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", errors.New(""))
	_, status := usecaseCat.InsertMuncipalities(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertMunicipalitiesErrFindById(t *testing.T) {
	cat := entity.Municipalities{
		Nombre: "dasdasd",
		Model: entity.Model{
			ID: 1,
		},
	}

	dep := entity.Department{
		Nombre: "dasdasd",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryMunicipalities(t)
	repoDep := mocks.NewIRepositoryDepartment(t)

	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewMunicipalitiesUseCase(repoCat, repoDep, objectFile)
	repoDep.On("FindDepartmentById", mock.Anything).Return(&dep, errors.New(""))
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.InsertMuncipalities(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_InsertMunicipalitiesErrInsert(t *testing.T) {
	cat := entity.Municipalities{
		Nombre: "dasdasd",
		Model: entity.Model{
			ID: 1,
		},
	}

	dep := entity.Department{
		Nombre: "dasdasd",
		Model: entity.Model{
			ID: 1,
		},
	}
	repoCat := mocks.NewIRepositoryMunicipalities(t)
	repoDep := mocks.NewIRepositoryDepartment(t)

	objectFile := mocks.NewIGCImageRepo(t)
	usecaseCat := usecase.NewMunicipalitiesUseCase(repoCat, repoDep, objectFile)
	repoCat.On("InsertMunicipalitiest", mock.Anything).Return(&cat, errors.New(""))
	repoDep.On("FindDepartmentById", mock.Anything).Return(&dep, nil)
	objectFile.On("SetFile", mock.Anything, mock.Anything, mock.Anything).Return("", nil)
	_, status := usecaseCat.InsertMuncipalities(context.Background(), cat)
	assert.Equal(t, status, http.StatusBadRequest)
}

func Test_DeleteMunicipalitiesOk(t *testing.T) {
	repoCat := mocks.NewIRepositoryMunicipalities(t)
	repoDep := mocks.NewIRepositoryDepartment(t)
	dep := entity.Department{
		Nombre:  "dasdasd",
		Bandera: "http://localjost.com/dasda",
		Model: entity.Model{
			ID: 1,
		},
	}
	cat := entity.Municipalities{
		Nombre:  "dasdasd",
		Bandera: "http://localjost.com/dasda",

		Model: entity.Model{
			ID: 1,
		},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	repoDep.On("FindDepartmentById", mock.Anything).Return(&dep, nil)
	repoCat.On("FindOneMunicipalities", mock.Anything).Return(&cat, nil)
	repoCat.On("DeleteMunicipalities", mock.Anything).Return(nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	usecaseCat := usecase.NewMunicipalitiesUseCase(repoCat, repoDep, objectFile)
	_, status := usecaseCat.DeleteMuncipalities(context.Background(), 1)
	assert.Equal(t, status, http.StatusOK)

}

func Test_DeleteMunicipalitiesErrFindDepartment(t *testing.T) {
	repoCat := mocks.NewIRepositoryMunicipalities(t)
	repoDep := mocks.NewIRepositoryDepartment(t)
	dep := entity.Department{
		Nombre:  "dasdasd",
		Bandera: "http://localjost.com/dasda",
		Model: entity.Model{
			ID: 1,
		},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	repoDep.On("FindDepartmentById", mock.Anything).Return(&dep, errors.New(""))
	usecaseCat := usecase.NewMunicipalitiesUseCase(repoCat, repoDep, objectFile)
	_, status := usecaseCat.DeleteMuncipalities(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)

}

func Test_DeleteMunicipalitiesErrFindMun(t *testing.T) {
	repoCat := mocks.NewIRepositoryMunicipalities(t)
	repoDep := mocks.NewIRepositoryDepartment(t)
	dep := entity.Department{
		Nombre:  "dasdasd",
		Bandera: "http://localjost.com/dasda",
		Model: entity.Model{
			ID: 1,
		},
	}
	cat := entity.Municipalities{
		Nombre:  "dasdasd",
		Bandera: "http://localjost.com/dasda",

		Model: entity.Model{
			ID: 1,
		},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	repoDep.On("FindDepartmentById", mock.Anything).Return(&dep, nil)
	repoCat.On("FindOneMunicipalities", mock.Anything).Return(&cat, errors.New(""))
	usecaseCat := usecase.NewMunicipalitiesUseCase(repoCat, repoDep, objectFile)
	_, status := usecaseCat.DeleteMuncipalities(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)

}

func Test_DeleteMunicipalitiesErrFile(t *testing.T) {
	repoCat := mocks.NewIRepositoryMunicipalities(t)
	repoDep := mocks.NewIRepositoryDepartment(t)
	dep := entity.Department{
		Nombre:  "dasdasd",
		Bandera: "http://localjost.com/dasda",
		Model: entity.Model{
			ID: 1,
		},
	}
	cat := entity.Municipalities{
		Nombre:  "dasdasd",
		Bandera: "http://localjost.com/dasda",

		Model: entity.Model{
			ID: 1,
		},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	repoDep.On("FindDepartmentById", mock.Anything).Return(&dep, nil)
	repoCat.On("FindOneMunicipalities", mock.Anything).Return(&cat, nil)
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(""))
	usecaseCat := usecase.NewMunicipalitiesUseCase(repoCat, repoDep, objectFile)
	_, status := usecaseCat.DeleteMuncipalities(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)

}

func Test_DeleteMunicipalitiesErrDelete(t *testing.T) {
	repoCat := mocks.NewIRepositoryMunicipalities(t)
	repoDep := mocks.NewIRepositoryDepartment(t)
	dep := entity.Department{
		Nombre:  "dasdasd",
		Bandera: "http://localjost.com/dasda",
		Model: entity.Model{
			ID: 1,
		},
	}
	cat := entity.Municipalities{
		Nombre:  "dasdasd",
		Bandera: "http://localjost.com/dasda",

		Model: entity.Model{
			ID: 1,
		},
	}
	objectFile := mocks.NewIGCImageRepo(t)
	repoDep.On("FindDepartmentById", mock.Anything).Return(&dep, nil)
	repoCat.On("FindOneMunicipalities", mock.Anything).Return(&cat, nil)
	repoCat.On("DeleteMunicipalities", mock.Anything).Return(errors.New(""))
	objectFile.On("DeleteFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	usecaseCat := usecase.NewMunicipalitiesUseCase(repoCat, repoDep, objectFile)
	_, status := usecaseCat.DeleteMuncipalities(context.Background(), 1)
	assert.Equal(t, status, http.StatusBadRequest)

}
