package handler_test

import (
	"service-sites/cmd/handler"
	"service-sites/internal/domain/usecase"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_HandlerCategory(t *testing.T) {
	e := handler.NewHandlerCategory(echo.New(), usecase.CategoryUseCase{}, func(next echo.HandlerFunc) echo.HandlerFunc { return nil })
	assert.NotNil(t, e)
}
