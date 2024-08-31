package model

import "time"

type Category struct {
	ID        int       `gorm:"autoIncrement;not null;primaryKey" json:"id"`
	Grouping  string    `gorm:"type:text" json:"grouping"`
	Name      string    `gorm:"type:text;not null" json:"name"`
	CreatedAt time.Time `gorm:"precision:6" json:"createdAt"`
	UpdatedAt time.Time `gorm:"precision:6" json:"updatedAt"`
}

func (*Category) TableName() string {
	return "categories"
}
