package validator_test

import (
	"net/http"
	"net/http/httptest"
	"service-sites/internal/domain/validator"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/labstack/echo/v4"
)

func TestValidateSubcategory_Fail(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"field": "value"}`))
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	nextCalled := false
	handler := func(c echo.Context) error {
		nextCalled = true
		return nil
	}

	validator.ValidateSubCategory(handler)(c)
	assert.Equal(t, !nextCalled, true)
}
