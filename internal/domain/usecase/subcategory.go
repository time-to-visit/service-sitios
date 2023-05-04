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

type SubCategoryUseCase struct {
	repoSubCate repository.IRepositorySubCategory
	repoCate    repository.IRepositoryCategory
	file        storage.IGCImageRepo
}

func NewSubCategoryUseCase(repoSubCate repository.IRepositorySubCategory, repoCate repository.IRepositoryCategory, file storage.IGCImageRepo) SubCategoryUseCase {
	return SubCategoryUseCase{
		repoSubCate,
		repoCate,
		file,
	}
}

func (e *SubCategoryUseCase) InsertSubcategory(ctx context.Context, subCategory entity.Subcategory) (interface{}, int) {
	category, err := e.repoCate.FindCategoryById(subCategory.CategoryId)
	if err != nil || category.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "problemas al buscar la categoria", nil), http.StatusBadRequest
	}
	pathname, err := e.file.SetFile(ctx, subCategory.UrlImagen, "subcategories/cat-%s.png")
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	subCategory.UrlImagen = pathname
	newSubcategory, err := e.repoSubCate.InsertSubcategory(subCategory)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar la subcategoria", nil), http.StatusBadRequest

	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "municipio insertada exitosamente", newSubcategory), http.StatusOK

}

func (e *SubCategoryUseCase) DeleteSubCategory(ctx context.Context, idCategory int, idSubcategory int) (interface{}, int) {
	subCategories, err := e.repoSubCate.FindOneSubcategory(idSubcategory)
	if err != nil || subCategories.ID == 0 {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "la category no fue encontrado", nil), http.StatusBadRequest
	}
	objectName := utils.ExtractObjectName(subCategories.UrlImagen)
	err = e.file.DeleteFile(ctx, "subcategories/%s", objectName)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema con storage", nil), http.StatusBadRequest
	}
	err = e.repoSubCate.DeleteSubcategory(idCategory, idSubcategory)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar la subcateoria", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "subcategoria eliminada exitosamente", nil), http.StatusOK

}
