package repository

import (
	"github.com/gofrs/uuid"
)

type GroupWithParticipantCount struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Currency         string    `json:"currency"`
	ParticipantCount int       `json:"participantCount"`
}

type GroupRepository interface {
	GetGroups(groupIds []uuid.UUID) ([]*GroupWithParticipantCount, error)
}
