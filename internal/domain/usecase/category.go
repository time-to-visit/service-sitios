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

type CategoryUseCase struct {
	repoCate repository.IRepositoryCategory
	file     storage.IGCImageRepo
}

func NewCategoryUseCase(repoCate repository.IRepositoryCategory, file storage.IGCImageRepo) CategoryUseCase {
	return CategoryUseCase{
		repoCate,
		file,
	}
}

func (e *CategoryUseCase) SeeCategory() (interface{}, int) {
	category, err := e.repoCate.FindCategory()
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la informacion", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "sucess", category), http.StatusOK
}

func (e *CategoryUseCase) UpdateCategory(category entity.Category) (interface{}, int) {
	categoryUpdate, err := e.repoCate.UpdateCategory(category)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al actualizar la categoria", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "categoria actualizada exitosamente", categoryUpdate), http.StatusOK
}

func (e *CategoryUseCase) DeleteCategory(ctx context.Context, Id int) (interface{}, int) {
	cat, err := e.repoCate.FindCategoryById(Id)
	if err != nil || cat.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "la category no fue encontrado", nil), http.StatusBadRequest
	}
	objectName := utils.ExtractObjectName(cat.UrlImagen)
	err = e.file.DeleteFile(ctx, "category/%s", objectName)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	err = e.repoCate.DeleteCategory(Id)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar la categoria", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "categoria eliminada exitosamente", nil), http.StatusOK

}

func (e *CategoryUseCase) InsertCategory(ctx context.Context, category entity.Category) (interface{}, int) {
	pathname, err := e.file.SetFile(ctx, category.UrlImagen, "category/cat-%s.png")
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	category.UrlImagen = pathname
	newCategory, err := e.repoCate.RegisterCategory(category)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar la categoria", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "categoria insertada exitosamente", newCategory), http.StatusOK
}
