package handler

import (
	"service-sites/cmd/entry"
	"service-sites/internal/domain/usecase"
	"service-sites/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerSites(e *echo.Echo, sitesCaseUse usecase.SitesUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	sitesEntry := entry.NewSitesEntry(sitesCaseUse)
	e.POST("/sites/sites", sitesEntry.AddSites, auth, validator.ValidateSites)
	e.GET("/sites/sites", sitesEntry.SeeSites, auth)
	e.DELETE("/sites/sites/:ID", sitesEntry.DeleteSites, auth)
	e.GET("/sites/sites/:ID", sitesEntry.SeeOneSites, auth)
	return e
}
