package extension

import (
	"net/http"

	"github.com/hellskater/udhaar-backend/internal/router/extension/herror"
	jsonIter "github.com/json-iterator/go"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// ErrorHandler Custom error handler
func ErrorHandler(logger *zap.Logger) echo.HTTPErrorHandler {
	return func(e error, c echo.Context) {
		var (
			code int
			body interface{}
		)

		switch err := e.(type) {
		case nil:
			return
		case *echo.HTTPError:
			if err.Internal != nil {
				if herr, ok := err.Internal.(*echo.HTTPError); ok {
					err = herr
				}
			}

			switch m := err.Message.(type) {
			case string:
				body = echo.Map{"message": m}
			case error:
				body = echo.Map{"message": m.Error()}
			default:
				body = echo.Map{"message": m}
			}

			code = err.Code
		case *herror.InternalError:
			logger.Error(err.Error(), append(err.Fields, zap.String("requestId", GetRequestID(c)))...)
			code = http.StatusInternalServerError
			body = echo.Map{"message": http.StatusText(http.StatusInternalServerError)}
		default:
			logger.Error(err.Error(), zap.String("requestId", GetRequestID(c)))
			code = http.StatusInternalServerError
			body = echo.Map{"message": http.StatusText(http.StatusInternalServerError)}
		}

		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				e = c.NoContent(code)
			} else {
				e = json(c, code, body, jsonIter.ConfigFastest)
			}
			if e != nil {
				logger.Warn("failed to send error response", zap.Error(e), zap.String("requestId", GetRequestID(c)))
			}
		}
	}
}
