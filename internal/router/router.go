package router

import (
	"compress/gzip"

	"github.com/hellskater/udhaar-backend/internal/repository"
	"github.com/hellskater/udhaar-backend/internal/router/extension"
	"github.com/hellskater/udhaar-backend/internal/router/middlewares"
	v1 "github.com/hellskater/udhaar-backend/internal/router/v1"
	"github.com/hellskater/udhaar-backend/internal/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Router struct {
	e  *echo.Echo
	v1 *v1.Handlers
}

func newRouter(db *gorm.DB, repo repository.Repository, ss *service.Services, logger *zap.Logger, config *Config) *Router {
	echo := newEcho(logger, config, repo)
	booksService := ss.Books
	handlers := &v1.Handlers{
		Books:  booksService,
		Repo:   repo,
		Logger: logger,
	}
	router := &Router{
		e:  echo,
		v1: handlers,
	}
	return router
}

func Setup(db *gorm.DB, repo repository.Repository, ss *service.Services, logger *zap.Logger, config *Config) *echo.Echo {
	r := newRouter(db, repo, ss, logger.Named("router"), config)

	api := r.e.Group("/api")
	r.v1.Setup(api)

	return r.e
}

func newEcho(logger *zap.Logger, config *Config, repo repository.Repository) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = extension.ErrorHandler(logger)

	// Middleware settings
	e.Use(middlewares.RequestID())
	e.Use(middlewares.AccessLogging(logger.Named("access_log"), config.Development))

	e.Use(middlewares.Recovery(logger))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: gzip.BestSpeed,
	}))

	e.Use(extension.Wrap(repo))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		ExposeHeaders: []string{echo.HeaderXRequestID},
		AllowHeaders:  []string{echo.HeaderContentType, echo.HeaderAuthorization},
		MaxAge:        3600,
	}))

	return e
}
