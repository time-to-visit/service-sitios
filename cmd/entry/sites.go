package entry

import (
	"service-sites/internal/domain/entity"
	"service-sites/internal/domain/usecase"
	"service-sites/internal/domain/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type sitesEntry struct {
	sitesCaseuse usecase.SitesUseCase
}

func NewSitesEntry(sitesCaseuse usecase.SitesUseCase) *sitesEntry {
	return &sitesEntry{
		sitesCaseuse,
	}
}

func (e *sitesEntry) SeeSites(c echo.Context) error {
	filter := utils.UrlValuesToMap(c.QueryParams())
	response, status := e.sitesCaseuse.SeeSites(filter)
	return c.JSON(status, response)
}

func (e *sitesEntry) SeeOneSites(c echo.Context) error {
	idSites, _ := strconv.Atoi(c.Param("ID"))
	response, status := e.sitesCaseuse.SeeOneSites(idSites)
	return c.JSON(status, response)
}

func (e *sitesEntry) DeleteSites(c echo.Context) error {
	idSites, _ := strconv.Atoi(c.Param("ID"))
	response, status := e.sitesCaseuse.DeleteSites(idSites)
	return c.JSON(status, response)
}

func (e *sitesEntry) AddSites(c echo.Context) error {
	sites := c.Get("sites").(*entity.Sites)
	response, status := e.sitesCaseuse.AddSites(*sites)
	return c.JSON(status, response)
}
