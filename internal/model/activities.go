package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type ActivityType string

const (
	UpdateGroup   ActivityType = "UPDATE_GROUP"
	CreateExpense ActivityType = "CREATE_EXPENSE"
	UpdateExpense ActivityType = "UPDATE_EXPENSE"
	DeleteExpense ActivityType = "DELETE_EXPENSE"
)

type Activity struct {
	ID            uuid.UUID    `gorm:"type:char(36);not null;primaryKey" json:"id"`
	ActivityType  ActivityType `gorm:"type:varchar(30);not null" json:"activityType"`
	ParticipantID uuid.UUID    `gorm:"type:char(36);not null" json:"participantId"`
	GroupID       uuid.UUID    `gorm:"type:char(36);not null" json:"groupId"`
	ExpenseID     uuid.UUID    `gorm:"type:char(36);not null" json:"expenseId"`
	Data          string       `gorm:"type:text" json:"data"`
	CreatedAt     time.Time    `gorm:"precision:6" json:"createdAt"`
	UpdatedAt     time.Time    `gorm:"precision:6" json:"updatedAt"`

	Group Group `gorm:"constraint:activity_group_fk,OnDelete:CASCADE" json:"group"`
}

func (*Activity) TableName() string {
	return "activities"
}
