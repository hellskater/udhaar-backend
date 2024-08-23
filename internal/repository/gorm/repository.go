package gorm

import (
	"github.com/hellskater/udhaar-backend/internal/migration"
	"github.com/hellskater/udhaar-backend/internal/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	logger *zap.Logger
}

// NewGormRepository Initialize and generate repository implementation。
// If the schema is initialized, return the init: true。
func NewGormRepository(db *gorm.DB, logger *zap.Logger, doMigration bool) (repo repository.Repository, init bool, err error) {
	repo = &Repository{
		db:     db,
		logger: logger.Named("repository"),
	}
	if doMigration {
		if init, err = migration.Migrate(db); err != nil {
			return nil, false, err
		}
	}
	return
}
