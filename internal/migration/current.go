package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/hellskater/udhaar-backend/internal/model"
)

// Migrations All database migration
//
// When doing a new migration, be sure to add it to the end of this array.
func Migrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{}
}

// AllTables all table model of the latest schema
//
// Write the model structure of all the latest schema tables
func AllTables() []interface{} {
	return []interface{}{
		&model.Group{},
	}
}
