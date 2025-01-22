package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite3Config struct {
	Path string
}

func newSqliteDB(c *Sqlite3Config) (*Handler, error) {
	db, err := gorm.Open(sqlite.Open(c.Path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Handler{db: db}, nil
}
