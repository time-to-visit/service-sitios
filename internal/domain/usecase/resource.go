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

type ResourceUseCase struct {
	repoRes repository.IRepositoryResource
	file    storage.IGCImageRepo
}

func NewResourceUseCase(repoRes repository.IRepositoryResource, file storage.IGCImageRepo) ResourceUseCase {
	return ResourceUseCase{
		repoRes,
		file,
	}
}

func (e *ResourceUseCase) AddResource(ctx context.Context, resource entity.Resource) (interface{}, int) {
	if resource.TypeRecursos == "IMAGE" {
		pathname, err := e.file.SetFile(ctx, resource.Payload, "resource-sites/res-%s.png")

		if err != nil {
			return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
		}
		resource.Payload = pathname
	}
	newResource, err := e.repoRes.InsertResource(resource)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar el recurso", nil), http.StatusBadRequest

	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "recurso insertada exitosamente", newResource), http.StatusOK
}

func (e *ResourceUseCase) DeleteResource(ctx context.Context, idSites int, idResource int) (interface{}, int) {
	res, err := e.repoRes.FindOneResource(idSites, idResource)
	if err != nil || res.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "el resource no fue encontrado", nil), http.StatusBadRequest
	}
	objectName := utils.ExtractObjectName(res.Payload)
	err = e.file.DeleteFile(ctx, "resource-sites/%s", objectName)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	err = e.repoRes.DeleteResource(idSites, idResource)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el recurso", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "recurso eliminado exitosamente", nil), http.StatusOK
}
