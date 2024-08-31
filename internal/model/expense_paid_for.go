package model

import "github.com/gofrs/uuid"

type ExpensePaidFor struct {
	ExpenseID     uuid.UUID `gorm:"type:char(36);not null;primaryKey" json:"expenseId"`
	ParticipantID uuid.UUID `gorm:"type:char(36);not null;primaryKey" json:"participantId"`
	Shares        int       `gorm:"type:int;not null;default:1" json:"shares"`

	Expense     Expense     `gorm:"constraint:expensepaidfor_expense_fk,OnDelete:CASCADE" json:"expense"`
	Participant Participant `gorm:"constraint:expensepaidfor_participant_fk,OnDelete:CASCADE" json:"participant"`
}

func (*ExpensePaidFor) TableName() string {
	return "expense_paid_for"
}
