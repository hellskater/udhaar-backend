package middlewares

import (
	"github.com/hellskater/udhaar-backend/internal/router/extension"
	"github.com/labstack/echo"
)

// RequestID Middleware that generates a request ID
func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderXRequestID, extension.GetRequestID(c))
			return next(c)
		}
	}
}
