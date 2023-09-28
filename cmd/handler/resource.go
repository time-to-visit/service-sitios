package handler

import (
	"service-sites/cmd/entry"
	"service-sites/internal/domain/usecase"
	"service-sites/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerResource(e *echo.Echo, ResourceCaseUse usecase.ResourceUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	resourceEntry := entry.NewResourceEntry(ResourceCaseUse)
	e.POST("/sites/resource", resourceEntry.AddResource, auth, validator.ValidateResource)
	e.DELETE("/sites/resource/:IDSITES/:IDRESOURCE", resourceEntry.DeleteResource, auth)
	return e
}
