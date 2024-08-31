package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type Participant struct {
	ID        uuid.UUID `gorm:"type:char(36);not null;primaryKey" json:"id"`
	Name      string    `gorm:"type:text;not null" json:"name"`
	GroupID   uuid.UUID `gorm:"type:char(36);not null" json:"groupId"`
	CreatedAt time.Time `gorm:"precision:6" json:"createdAt"`
	UpdatedAt time.Time `gorm:"precision:6" json:"updatedAt"`

	Group           Group             `gorm:"constraint:participant_group_fk,OnDelete:CASCADE" json:"-"`
	ExpensesPaidBy  []*Expense        `gorm:"constraint:participant_expensespaidby_fk,OnDelete:CASCADE;foreignKey:PaidByID" json:"-"`
	ExpensesPaidFor []*ExpensePaidFor `gorm:"constraint:participant_expensespaidfor_fk,OnDelete:CASCADE" json:"-"`
}

func (*Participant) TableName() string {
	return "participants"
}
