//go:build wireinject
// +build wireinject

package router

import (
	"github.com/google/wire"
	"github.com/hellskater/udhaar-backend/internal/repository"
	v1 "github.com/hellskater/udhaar-backend/internal/router/v1"
	"github.com/hellskater/udhaar-backend/internal/service"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func newRouter(db *gorm.DB, repo repository.Repository, ss *service.Services, logger *zap.Logger, config *Config) *Router {
	wire.Build(
		service.ProviderSet,
		newEcho,
		wire.Struct(new(v1.Handlers), "*"),
		wire.Struct(new(Router), "*"),
	)
	return nil
}
