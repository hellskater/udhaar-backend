package gorm

import (
	"github.com/hellskater/udhaar-backend/internal/repository"
	"gorm.io/gorm"
)

func convertError(err error) error {
	switch {
	case err == gorm.ErrRecordNotFound:
		return repository.ErrNotFound
	default:
		return err
	}
}
