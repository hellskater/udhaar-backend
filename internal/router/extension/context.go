package extension

import (
	vd "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofrs/uuid"
	"github.com/hellskater/udhaar-backend/internal/repository"
	"github.com/hellskater/udhaar-backend/internal/router/consts"
	"github.com/hellskater/udhaar-backend/internal/router/extension/herror"
	"github.com/hellskater/udhaar-backend/internal/router/utils"
	jsonIter "github.com/json-iterator/go"
	"github.com/labstack/echo"
)

// Context Echo.context custom
type Context struct {
	echo.Context
}

// JSON Replace Encoding/json with jsonitor.configCompatiblewithstandardlibrary
func (c *Context) JSON(code int, i interface{}) (err error) {
	if _, pretty := c.QueryParams()["pretty"]; pretty {
		return c.Context.JSON(code, i)
	}
	return json(c, code, i, jsonIter.ConfigFastest)
}

func json(c echo.Context, code int, i interface{}, cfg jsonIter.API) error {
	stream := cfg.BorrowStream(c.Response())
	defer cfg.ReturnStream(stream)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(code)
	stream.WriteVal(i)
	stream.WriteRaw("\n")
	return stream.Flush()
}

// Wrap Custom Context wrapper
func Wrap(repo repository.Repository) echo.MiddlewareFunc {
	return func(n echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(consts.KeyRepo, repo)
			return n(&Context{Context: c})
		}
	}
}

// GetRequestParamAsUUID Get the specified request parameter as UUID
func GetRequestParamAsUUID(c echo.Context, name string) uuid.UUID {
	return uuid.FromStringOrNil(c.Param(name))
}

// BindAndValidate Decoritize FormData or JSON to the structure I
func BindAndValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	if err := vd.ValidateWithContext(utils.NewRequestValidateContext(c), i); err != nil {
		if e, ok := err.(vd.InternalError); ok {
			return herror.InternalServerError(e.InternalError())
		}
		return herror.BadRequest(err)
	}
	return nil
}
