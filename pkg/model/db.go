package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Config struct {
	Type    string
	Sqlite3 *Sqlite3Config
}

type Database interface {
	Create() error
	Close()
}

type Handler struct {
	db *gorm.DB
}

func New(c *Config) (*Handler, error) {
	if c.Type == "sqlite3" {
		return newSqliteDB(c.Sqlite3)
	}
	return nil, fmt.Errorf("not supported database type: %s", c.Type)
}
