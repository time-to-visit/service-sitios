package usecase

import (
	"context"
	"net/http"
	"service-sites/internal/domain/entity"
	objectValues "service-sites/internal/domain/object_values"
	"service-sites/internal/domain/repository"
	"service-sites/internal/infra/storage"
	"service-sites/internal/utils"
)

type DepartmentUseCase struct {
	repoDep repository.IRepositoryDepartment
	file    storage.IGCImageRepo
}

func NewDepartmentUseCase(repoDep repository.IRepositoryDepartment, file storage.IGCImageRepo) DepartmentUseCase {
	return DepartmentUseCase{
		repoDep,
		file,
	}
}

func (e *DepartmentUseCase) SeeDeparment() (interface{}, int) {
	deparment, err := e.repoDep.FindDepartment()
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", deparment), http.StatusOK
}

func (e *DepartmentUseCase) InsertDeparment(ctx context.Context, deparment entity.Department) (interface{}, int) {
	pathname, err := e.file.SetFile(ctx, deparment.Bandera, "deparment/dep-%s.png")
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	deparment.Bandera = pathname
	newDepartment, err := e.repoDep.InsertDepartment(deparment)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar el departamento", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "departamento insertada exitosamente", newDepartment), http.StatusOK
}

func (e *DepartmentUseCase) DeleteDeparment(ctx context.Context, id int) (interface{}, int) {
	dep, err := e.repoDep.FindDepartmentById(id)
	if err != nil || dep == nil || dep.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "no se encontro el steps", nil), http.StatusBadRequest
	}
	objectName := utils.ExtractObjectName(dep.Bandera)
	err = e.file.DeleteFile(ctx, "deparment/%s", objectName)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	err = e.repoDep.DeleteDepartment(id)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el departamento", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "departamento eliminada exitosamente", nil), http.StatusOK
}
