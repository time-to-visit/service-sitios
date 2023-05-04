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

type MunicipalitiesUseCase struct {
	repoMun repository.IRepositoryMunicipalities
	repoDep repository.IRepositoryDepartment
	file    storage.IGCImageRepo
}

func NewMunicipalitiesUseCase(repoMun repository.IRepositoryMunicipalities, repoDep repository.IRepositoryDepartment, file storage.IGCImageRepo) MunicipalitiesUseCase {
	return MunicipalitiesUseCase{
		repoMun,
		repoDep,
		file,
	}
}

func (e *MunicipalitiesUseCase) SeeMuncipalities(idDeparment int) (interface{}, int) {
	municipalities, err := e.repoMun.FindMunicipalities(idDeparment)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", municipalities), http.StatusOK
}

func (e *MunicipalitiesUseCase) InsertMuncipalities(ctx context.Context, municipalities entity.Municipalities) (interface{}, int) {
	pathname, err := e.file.SetFile(ctx, municipalities.Bandera, "municipalities/step-%s.png")
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	municipalities.Bandera = pathname
	deparment, err := e.repoDep.FindDepartmentById(municipalities.DepartmentId)
	if err != nil || deparment.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "el departamento no fue encontrado", nil), http.StatusBadRequest
	}

	newMunicipalities, err := e.repoMun.InsertMunicipalitiest(municipalities)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar el municipio", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "municipio insertada exitosamente", newMunicipalities), http.StatusOK

}

func (e *MunicipalitiesUseCase) DeleteMuncipalities(ctx context.Context, Id int) (interface{}, int) {
	deparment, err := e.repoDep.FindDepartmentById(Id)
	if err != nil || deparment.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "el departament no fue encontrado", nil), http.StatusBadRequest
	}
	mun, err := e.repoMun.FindOneMunicipalities(Id)
	if err != nil || mun == nil || mun.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "el municipalities no fue encontrado", nil), http.StatusBadRequest
	}
	objectName := utils.ExtractObjectName(mun.Bandera)
	err = e.file.DeleteFile(ctx, "municipalities/%s", objectName)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	errDel := e.repoMun.DeleteMunicipalities(Id)
	if errDel != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el municipio", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "municipio eliminado exitosamente", nil), http.StatusOK
}
