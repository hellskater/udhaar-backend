package herror

import (
	"net/http"

	"github.com/hellskater/udhaar-backend/internal/repository"
	"github.com/labstack/echo"
)

func NotFound(err ...interface{}) error {
	return HTTPError(http.StatusNotFound, err)
}

func BadRequest(err ...interface{}) error {
	return HTTPError(http.StatusBadRequest, err)
}

func Forbidden(err ...interface{}) error {
	return HTTPError(http.StatusForbidden, err)
}

func Conflict(err ...interface{}) error {
	return HTTPError(http.StatusConflict, err)
}

func Unauthorized(err ...interface{}) error {
	return HTTPError(http.StatusUnauthorized, err)
}

func HTTPError(code int, err interface{}) error {
	switch v := err.(type) {
	case []interface{}:
		if len(v) > 0 {
			return HTTPError(code, v[0])
		}
		return HTTPError(code, nil)
	case string:
		return echo.NewHTTPError(code, v)
	case *repository.ArgumentError:
		return echo.NewHTTPError(code, v.Error())
	case nil:
		return echo.NewHTTPError(code)
	default:
		return echo.NewHTTPError(code, v)
	}
}
