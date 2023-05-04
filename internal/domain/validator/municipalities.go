package validator

import (
	"net/http"
	"service-sites/internal/domain/entity"
	validatorPer "service-sites/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateMunicipalities(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		municipalities := new(entity.Municipalities)

		_ = c.Bind(&municipalities)
		if err := v.Struct(municipalities); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("municipalities", municipalities)
		return next(c)
	}
}
