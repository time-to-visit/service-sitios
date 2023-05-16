package handler_test

import (
	"service-sites/cmd/handler"
	"service-sites/internal/domain/usecase"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_HandlerSubcategory(t *testing.T) {
	e := handler.NewHandlerSubCategory(echo.New(), usecase.SubCategoryUseCase{}, func(next echo.HandlerFunc) echo.HandlerFunc { return nil })
	assert.NotNil(t, e)
}
