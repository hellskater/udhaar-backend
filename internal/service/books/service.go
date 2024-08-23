package books

import (
	"github.com/hellskater/udhaar-backend/internal/repository"
	"go.uber.org/zap"
)

type serviceImpl struct {
	repo   repository.Repository
	logger *zap.Logger
}

type Service interface{}

func NewService(repo repository.Repository, logger *zap.Logger) Service {
	p := &serviceImpl{
		repo:   repo,
		logger: logger,
	}

	return p
}
