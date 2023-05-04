package entry

import (
	"service-sites/internal/domain/entity"
	"service-sites/internal/domain/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type reviewEntry struct {
	reviewCaseuse usecase.ReviewUseCase
}

func NewReviewEntry(reviewCaseuse usecase.ReviewUseCase) *reviewEntry {
	return &reviewEntry{
		reviewCaseuse,
	}
}

func (e *reviewEntry) AddReview(c echo.Context) error {
	review := c.Get("review").(*entity.Review)
	response, status := e.reviewCaseuse.AddReview(*review)
	return c.JSON(status, response)
}

func (e *reviewEntry) DeleteReview(c echo.Context) error {
	idSites, _ := strconv.Atoi(c.Param("IDSITES"))
	idReview, _ := strconv.Atoi(c.Param("IDREVIEW"))
	response, status := e.reviewCaseuse.DeleteReview(idSites, idReview)
	return c.JSON(status, response)
}

func (e *reviewEntry) SeeReview(c echo.Context) error {
	idSites, _ := strconv.Atoi(c.Param("ID"))
	response, status := e.reviewCaseuse.SeeReview(idSites)
	return c.JSON(status, response)
}
