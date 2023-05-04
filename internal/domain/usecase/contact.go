package usecase

import (
	"net/http"
	"service-sites/internal/domain/entity"
	objectValues "service-sites/internal/domain/object_values"
	"service-sites/internal/domain/repository"
)

type ContactUseCase struct {
	repoCon repository.IRepositoryContact
}

func NewContactUseCase(repoCon repository.IRepositoryContact) ContactUseCase {
	return ContactUseCase{
		repoCon,
	}
}

func (e *ContactUseCase) DeleteContact(IdSites int, IdContact int) (interface{}, int) {
	err := e.repoCon.DeleteContact(IdSites, IdContact)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al eliminar el contacto", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "contacto eliminada exitosamente", nil), http.StatusOK
}

func (e *ContactUseCase) InsertContact(contact entity.Contact) (interface{}, int) {
	newContact, err := e.repoCon.InsertContact(contact)
	if err != nil {
		return objectValues.NewResponseWithData(http.StatusBadRequest, "error", "hubo un problema al insertar el contacto", nil), http.StatusBadRequest
	}
	return objectValues.NewResponseWithData(http.StatusOK, "ok", "contacto insertada exitosamente", newContact), http.StatusOK
}
