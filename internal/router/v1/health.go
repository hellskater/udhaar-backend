package v1

import (
	"net/http"

	"github.com/labstack/echo"
)

// Health GET /health
func (h *Handlers) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
