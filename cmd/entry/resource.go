package entry

import (
	"service-sites/internal/domain/entity"
	"service-sites/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type resourceEntry struct {
	resourceCaseuse usecase.ResourceUseCase
}

func NewResourceEntry(resourceCaseuse usecase.ResourceUseCase) *resourceEntry {
	return &resourceEntry{
		resourceCaseuse,
	}
}

func (e *resourceEntry) AddResource(c echo.Context) error {
	resource := c.Get("resource").(*entity.Resource)
	response, status := e.resourceCaseuse.AddResource(c.Request().Context(), *resource)
	return c.JSON(status, response)
}

func (e *resourceEntry) DeleteResource(c echo.Context) error {
	idSites, _ := strconv.Atoi(c.Param("IDSITES"))
	idResource, _ := strconv.Atoi(c.Param("IDRESOURCE"))
	response, status := e.resourceCaseuse.DeleteResource(c.Request().Context(), idSites, idResource)
	return c.JSON(status, response)
}
