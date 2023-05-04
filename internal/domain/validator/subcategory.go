package validator

import (
	"net/http"
	"service-sites/internal/domain/entity"
	validatorPer "service-sites/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateSubCategory(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		subCategory := new(entity.Subcategory)

		_ = c.Bind(&subCategory)
		if err := v.Struct(subCategory); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("subCategory", subCategory)
		return next(c)
	}
}
