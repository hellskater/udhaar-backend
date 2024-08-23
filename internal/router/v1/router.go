package v1

import (
	"github.com/hellskater/udhaar-backend/internal/repository"
	"github.com/hellskater/udhaar-backend/internal/service/books"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type Handlers struct {
	Books  books.Service
	Repo   repository.Repository
	Logger *zap.Logger
}

func (h *Handlers) Setup(e *echo.Group) {
	// standard
	api := e.Group("/v1")
	{
		api.GET("/health", h.Health)

		// groups
		apiGroups := api.Group("/groups")
		{
			apiGroups.GET("", h.GetGroups)
		}
	}
}
