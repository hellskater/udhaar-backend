package gorm

import "github.com/hellskater/udhaar-backend/internal/model"

func (repo *Repository) GetGroups() ([]*model.Group, error) {
	groups := make([]*model.Group, 0)

	if err := repo.db.Find(&groups).Error; err != nil {
		return nil, convertError(err)
	}

	return groups, nil
}
