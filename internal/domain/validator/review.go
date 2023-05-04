package validator

import (
	"fmt"
	"net/http"
	"service-sites/internal/domain/entity"
	objectValues "service-sites/internal/domain/object_values"
	validatorPer "service-sites/internal/infra/validation"

	validatorV "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateReview(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validatorPer.NewValidator()
		auth := c.Get("auth").(objectValues.Auth)
		review := new(entity.Review)
		_ = c.Bind(&review)
		if err := v.Struct(review); err != nil {
			errs := err.(validatorV.ValidationErrors)
			return c.JSON(http.StatusBadRequest, validatorPer.GenerateMessage(v, errs))
		}
		fmt.Println(auth)
		review.UserId = int(auth.Data.ID)
		review.UserNombre = auth.Data.Nombres
		c.Set("review", review)
		return next(c)
	}
}
