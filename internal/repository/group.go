package repository

import "github.com/hellskater/udhaar-backend/internal/model"

type GroupRepository interface {
	GetGroups() ([]*model.Group, error)
}
