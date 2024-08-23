//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"github.com/hellskater/udhaar-backend/internal/repository"
	"github.com/hellskater/udhaar-backend/internal/router"
	"github.com/hellskater/udhaar-backend/internal/service"
	"github.com/hellskater/udhaar-backend/internal/service/books"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func newServer(db *gorm.DB, repo repository.Repository, logger *zap.Logger, c *Config) (*Server, error) {
	wire.Build(
		books.NewService,
		router.Setup,
		provideRouterConfig,
		wire.Struct(new(service.Services), "*"),
		wire.Struct(new(Server), "*"),
	)
	return nil, nil
}
