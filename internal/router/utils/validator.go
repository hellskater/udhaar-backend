package utils

import (
	"context"

	"github.com/hellskater/udhaar-backend/internal/router/consts"
	"github.com/labstack/echo"
)

type ctxKey int

const (
	repoctxKey ctxKey = iota
)

func NewRequestValidateContext(c echo.Context) context.Context {
	return context.WithValue(context.Background(), repoctxKey, c.Get(consts.KeyRepo))
}
