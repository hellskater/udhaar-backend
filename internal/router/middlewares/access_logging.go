package middlewares

import (
	"strconv"
	"time"

	"github.com/blendle/zapdriver"
	"github.com/hellskater/udhaar-backend/internal/router/extension"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// AccessLogging middleware
func AccessLogging(logger *zap.Logger, dev bool) echo.MiddlewareFunc {
	if dev {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				start := time.Now()
				if err := next(c); err != nil {
					c.Error(err)
				}
				stop := time.Now()

				req := c.Request()
				res := c.Response()
				logger.Sugar().Infof("%3d | %s | %s %s %d", res.Status, stop.Sub(start), req.Method, req.URL, res.Size)
				return nil
			}
		}
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			req := c.Request()
			res := c.Response()
			logger.Info("",
				zap.String("requestId", extension.GetRequestID(c)),
				zapdriver.HTTP(&zapdriver.HTTPPayload{
					RequestMethod: req.Method,
					Status:        res.Status,
					UserAgent:     req.UserAgent(),
					RemoteIP:      c.RealIP(),
					Referer:       req.Referer(),
					Protocol:      req.Proto,
					RequestURL:    req.URL.String(),
					RequestSize:   req.Header.Get(echo.HeaderContentLength),
					ResponseSize:  strconv.FormatInt(res.Size, 10),
					Latency:       strconv.FormatFloat(stop.Sub(start).Seconds(), 'f', 9, 64) + "s",
				}))
			return nil
		}
	}
}
