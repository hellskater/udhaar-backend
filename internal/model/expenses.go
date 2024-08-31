package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type SplitMode string

const (
	Evenly       SplitMode = "EVENLY"
	ByShares     SplitMode = "BY_SHARES"
	ByPercentage SplitMode = "BY_PERCENTAGE"
	ByAmount     SplitMode = "BY_AMOUNT"
)

type Expense struct {
	ID          uuid.UUID `gorm:"type:char(36);not null;primaryKey" json:"id"`
	Title       string    `gorm:"type:text;not null" json:"title"`
	GroupID     uuid.UUID `gorm:"type:char(36);not null" json:"groupId"`
	ExpenseDate time.Time `gorm:"type:date;default:CURRENT_DATE" json:"expenseDate"`
	CategoryID  *int      `gorm:"type:int" json:"categoryId"`
	Amount      int       `gorm:"type:int;not null" json:"amount"`
	PaidByID    uuid.UUID `gorm:"type:char(36);not null" json:"paidById"`
	IsRepayment bool      `gorm:"type:boolean;default:false" json:"isRepayment"`
	SplitMode   SplitMode `gorm:"type:text;not null;default:'EVENLY'" json:"splitMode"`
	CreatedAt   time.Time `gorm:"precision:6" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"precision:6" json:"updatedAt"`
	Notes       string    `gorm:"type:text" json:"notes"`

	Group    Group       `gorm:"constraint:expense_group_fk,OnDelete:CASCADE" json:"group"`
	Category *Category   `gorm:"constraint:expense_category_fk,OnDelete:SET NULL" json:"category"`
	PaidBy   Participant `gorm:"constraint:expense_paidby_fk,OnDelete:CASCADE" json:"paidBy"`
}

func (*Expense) TableName() string {
	return "expenses"
}
