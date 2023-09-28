package handler

import (
	"service-sites/cmd/entry"
	"service-sites/internal/domain/usecase"
	"service-sites/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerMunicipalities(e *echo.Echo, MunicipalitiesCaseUse usecase.MunicipalitiesUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	municipalitiesEntry := entry.NewMunicipalitiesEntry(MunicipalitiesCaseUse)
	e.GET("/sites/municipalities/:ID", municipalitiesEntry.SeeMuncipalities, auth)
	e.POST("/sites/municipalities", municipalitiesEntry.InsertMuncipalities, auth, validator.ValidateMunicipalities)
	e.DELETE("/sites/municipalities/:ID", municipalitiesEntry.DeleteMuncipalities, auth)
	return e
}
