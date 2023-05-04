package config

import (
	"flag"
	"service-sites/cmd/handler"
	"service-sites/internal/domain/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	var configPath = ""
	configPath = *flag.String("config", "", "")

	if configPath == "" {
		configPath = "./data/config.yml"
	}

	setConfiguration(configPath)
}

func setConfiguration(configPath string) {
	Setup(configPath)
}

func Run(configPath string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	conf := GetConfig()
	setupDB(conf)
	ioc := genIoc(conf)

	e = handler.NewHandlerCategory(e, ioc["category"].(usecase.CategoryUseCase), AuthVerify)
	e = handler.NewHandlerContact(e, ioc["contact"].(usecase.ContactUseCase), AuthVerify)
	e = handler.NewHandlerDepartment(e, ioc["department"].(usecase.DepartmentUseCase), AuthVerify)
	e = handler.NewHandlerMunicipalities(e, ioc["municipalities"].(usecase.MunicipalitiesUseCase), AuthVerify)
	e = handler.NewHandlerResource(e, ioc["resource"].(usecase.ResourceUseCase), AuthVerify)
	e = handler.NewHandlerReview(e, ioc["review"].(usecase.ReviewUseCase), AuthVerify)
	e = handler.NewHandlerSites(e, ioc["sites"].(usecase.SitesUseCase), AuthVerify)
	e = handler.NewHandlerSubCategory(e, ioc["subcategory"].(usecase.SubCategoryUseCase), AuthVerify)
	e.Logger.Fatal(e.Start(":" + conf.Server.Port))
}
