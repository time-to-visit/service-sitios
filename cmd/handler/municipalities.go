package handler

import (
	"service-sites/cmd/entry"
	"service-sites/internal/domain/usecase"
	"service-sites/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerMunicipalities(e *echo.Echo, MunicipalitiesCaseUse usecase.MunicipalitiesUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	municipalitiesEntry := entry.NewMunicipalitiesEntry(MunicipalitiesCaseUse)
	e.GET("/municipalities/:ID", municipalitiesEntry.SeeMuncipalities, auth)
	e.POST("/municipalities", municipalitiesEntry.InsertMuncipalities, auth, validator.ValidateMunicipalities)
	e.DELETE("/municipalities/:ID", municipalitiesEntry.DeleteMuncipalities, auth)
	return e
}
