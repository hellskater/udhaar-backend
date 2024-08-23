package migration

import (
	"database/sql"
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Migrate: Execute the database migration
// If the schema is initialized in the first execution, return True with INIT.
func Migrate(db *gorm.DB) (init bool, err error) {
	m := gormigrate.New(db, &gormigrate.Options{
		TableName:                 "migrations",
		IDColumnName:              "id",
		IDColumnSize:              190,
		UseTransaction:            false,
		ValidateUnknownMigrations: true,
	}, Migrations())
	m.InitSchema(func(db *gorm.DB) error {
		// Called only for the first time
		// Write all the latest database definitions
		init = true

		// table
		return db.AutoMigrate(AllTables()...)
	})
	db.AutoMigrate(AllTables()...)
	err = m.Migrate()
	return
}

// DropAll Delete all database tables
func DropAll(db *gorm.DB) error {
	if err := db.Migrator().DropTable(AllTables()...); err != nil {
		return err
	}
	return db.Migrator().DropTable("migrations")
}

// CreateDatabasesIfNotExists Create if there is no database
func CreateDatabasesIfNotExists(dialect, dsn, prefix string, names ...string) error {
	conn, err := sql.Open(dialect, dsn)
	if err != nil {
		return err
	}
	defer conn.Close()
	for _, v := range names {
		_, err = conn.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s%s`", prefix, v))
		if err != nil {
			return err
		}
	}
	return nil
}
