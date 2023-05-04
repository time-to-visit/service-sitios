package handler

import (
	"service-sites/cmd/entry"
	"service-sites/internal/domain/usecase"
	"service-sites/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerContact(e *echo.Echo, ContactCaseUse usecase.ContactUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	contactEntry := entry.NewConcatEntry(ContactCaseUse)
	e.POST("/contact", contactEntry.InsertContact, auth, validator.ValidateContact)
	e.DELETE("/contact/:IDSITES/:IDCONTACT", contactEntry.DeleteContact, auth)
	return e
}
