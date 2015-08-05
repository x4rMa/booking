package database

import (
	"github.com/jinzhu/gorm"
)

type Database interface {
	DB() (*gorm.DB, error)
}
