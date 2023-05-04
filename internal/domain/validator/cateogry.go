package validator

import (
	"net/http"
	"service-sites/internal/domain/entity"
	validatorPer "service-sites/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateCategory(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		category := new(entity.Category)

		_ = c.Bind(&category)
		if err := v.Struct(category); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("category", category)
		return next(c)
	}
}
