package entry

import (
	"service-sites/internal/domain/entity"
	"service-sites/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type municipalitiesEntry struct {
	municipalitiesCaseuse usecase.MunicipalitiesUseCase
}

func NewMunicipalitiesEntry(municipalitiesCaseuse usecase.MunicipalitiesUseCase) *municipalitiesEntry {
	return &municipalitiesEntry{
		municipalitiesCaseuse,
	}
}

func (e *municipalitiesEntry) SeeMuncipalities(c echo.Context) error {
	id := c.Param("ID")
	idDepartment, _ := strconv.Atoi(id)
	response, status := e.municipalitiesCaseuse.SeeMuncipalities(idDepartment)
	return c.JSON(status, response)
}

func (e *municipalitiesEntry) InsertMuncipalities(c echo.Context) error {
	muncipalities := c.Get("municipalities").(*entity.Municipalities)
	response, status := e.municipalitiesCaseuse.InsertMuncipalities(c.Request().Context(), *muncipalities)
	return c.JSON(status, response)
}

func (e *municipalitiesEntry) DeleteMuncipalities(c echo.Context) error {
	id := c.Param("ID")
	idMunicipalities, _ := strconv.Atoi(id)
	response, status := e.municipalitiesCaseuse.DeleteMuncipalities(c.Request().Context(), idMunicipalities)
	return c.JSON(status, response)
}
