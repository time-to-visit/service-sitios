package handler

import (
	"service-sites/cmd/entry"
	"service-sites/internal/domain/usecase"
	"service-sites/internal/domain/validator"

	"github.com/labstack/echo/v4"
)

func NewHandlerReview(e *echo.Echo, ReviewCaseUse usecase.ReviewUseCase, auth func(next echo.HandlerFunc) echo.HandlerFunc) *echo.Echo {
	reviewEntry := entry.NewReviewEntry(ReviewCaseUse)
	e.GET("/review/:ID", reviewEntry.SeeReview, auth)
	e.DELETE("/review/:IDSITES/:IDREVIEW", reviewEntry.DeleteReview, auth)
	e.POST("/review", reviewEntry.AddReview, auth, validator.ValidateReview)
	return e
}
