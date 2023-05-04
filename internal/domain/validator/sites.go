package validator

import (
	"net/http"
	"service-sites/internal/domain/entity"
	validatorPer "service-sites/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateSites(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		sites := new(entity.Sites)

		_ = c.Bind(&sites)
		if err := v.Struct(sites); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("sites", sites)
		return next(c)
	}
}
