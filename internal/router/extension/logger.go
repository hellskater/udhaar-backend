package extension

import (
	"github.com/hellskater/udhaar-backend/pkg/utils/random"
	"github.com/labstack/echo"
)

// GetRequestID Return your request ID
func GetRequestID(c echo.Context) string {
	rid := c.Request().Header.Get(echo.HeaderXRequestID)
	if len(rid) == 0 {
		rid = random.AlphaNumeric(32)
	}
	return rid
}
