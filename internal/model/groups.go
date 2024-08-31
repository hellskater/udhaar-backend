package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type Group struct {
	ID        uuid.UUID `gorm:"type:char(36);not null;primaryKey" json:"id"`
	Name      string    `gorm:"type:text;not null" json:"name"`
	Currency  string    `gorm:"type:text;not null;default:'₹'" json:"currency"`
	CreatedAt time.Time `gorm:"precision:6" json:"createdAt"`
	UpdatedAt time.Time `gorm:"precision:6" json:"updatedAt"`

	Participants []*Participant `gorm:"constraint:group_participants_fk,OnDelete:SET NULL" json:"participants"`
	Expenses     []*Expense     `gorm:"constraint:group_expenses_fk,OnDelete:SET NULL" json:"expenses"`
	Activities   []*Activity    `gorm:"constraint:group_activities_fk,OnDelete:SET NULL" json:"activities"`
}

func (*Group) TableName() string {
	return "groups"
}
