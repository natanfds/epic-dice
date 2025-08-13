package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func MigrateSQLDB(db *gorm.DB, models []interface{}) error {
	return db.AutoMigrate(models...)
}

func CreateSQLDB(models ...interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("epic-dice.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if len(models) == 0 {
		return db, nil
	}
	if err := MigrateSQLDB(db, models); err != nil {
		return nil, err
	}
	return db, nil
}
