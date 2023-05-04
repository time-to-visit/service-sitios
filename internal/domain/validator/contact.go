package validator

import (
	"net/http"
	"service-sites/internal/domain/entity"
	validatorPer "service-sites/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateContact(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		contact := new(entity.Contact)

		_ = c.Bind(&contact)
		if err := v.Struct(contact); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		c.Set("contact", contact)
		return next(c)
	}
}
