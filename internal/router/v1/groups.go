package v1

import (
	"net/http"

	"github.com/hellskater/udhaar-backend/internal/router/extension/herror"
	"github.com/labstack/echo"
)

// GetGroups GET /groups
func (h *Handlers) GetGroups(c echo.Context) error {
	cfs, err := h.Repo.GetGroups()
	if err != nil {
		return herror.InternalServerError(err)
	}

	return c.JSON(http.StatusOK, cfs)
}
