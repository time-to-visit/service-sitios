package entry

import (
	"service-sites/internal/domain/entity"
	"service-sites/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type contactEntry struct {
	contactCaseuse usecase.ContactUseCase
}

func NewConcatEntry(contactCaseuse usecase.ContactUseCase) *contactEntry {
	return &contactEntry{
		contactCaseuse,
	}
}

func (e *contactEntry) DeleteContact(c echo.Context) error {
	c.Param("ID")
	idSites, _ := strconv.Atoi(c.Param("IDSITES"))
	idContact, _ := strconv.Atoi(c.Param("IDCONTACT"))
	response, status := e.contactCaseuse.DeleteContact(idSites, idContact)
	return c.JSON(status, response)
}

func (e *contactEntry) InsertContact(c echo.Context) error {
	contact := c.Get("contact").(*entity.Contact)
	response, status := e.contactCaseuse.InsertContact(*contact)
	return c.JSON(status, response)
}
