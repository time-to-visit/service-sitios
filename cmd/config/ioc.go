package config

import (
	"service-sites/internal/domain/repository"
	"service-sites/internal/domain/usecase"
	"service-sites/internal/infra/storage"
)

func genIoc(config *Configuration) map[string]interface{} {
	ioc := make(map[string]interface{})
	//database
	db := GetDB()
	repoCategory := repository.NewRepositoryCategory(db)
	repoContact := repository.NewRepositoryContact(db)
	repoDepartment := repository.NewRepositoryDepartment(db)
	repoMunicipalities := repository.NewRepositoryMunicipalities(db)
	repoResource := repository.NewRepositoryResource(db)
	repoReview := repository.NewRepositoryReview(db)
	repoSites := repository.NewRepositorySites(db)
	repoSubcategory := repository.NewRepositorySubCategory(db)
	clientStorage := storage.InitStorage(GetStorageClient(), config.Credential.Gcbucket)
	//ioc
	ioc["category"] = usecase.NewCategoryUseCase(repoCategory, clientStorage)
	ioc["contact"] = usecase.NewContactUseCase(repoContact)
	ioc["department"] = usecase.NewDepartmentUseCase(repoDepartment, clientStorage)
	ioc["municipalities"] = usecase.NewMunicipalitiesUseCase(repoMunicipalities, repoDepartment, clientStorage)
	ioc["resource"] = usecase.NewResourceUseCase(repoResource, clientStorage)
	ioc["review"] = usecase.NewReviewUseCase(repoReview, repoSites)
	ioc["sites"] = usecase.NewSitesUseCase(repoSites)
	ioc["subcategory"] = usecase.NewSubCategoryUseCase(repoSubcategory, repoCategory, clientStorage)

	return ioc
}
