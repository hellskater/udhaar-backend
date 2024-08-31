package gorm

import (
	"github.com/gofrs/uuid"
	"github.com/hellskater/udhaar-backend/internal/repository"
)

func (repo *Repository) GetGroups(groupIds []uuid.UUID) ([]*repository.GroupWithParticipantCount, error) {
	results := make([]*repository.GroupWithParticipantCount, 0)

	if err := repo.db.Table("groups").
		Select("groups.id, groups.name, groups.currency, COUNT(participants.id) as participant_count").
		Joins("left join participants on participants.group_id = groups.id").
		Where("groups.id IN (?)", groupIds).
		Group("groups.id").
		Find(&results).Error; err != nil {
		return nil, convertError(err)
	}

	return results, nil
}