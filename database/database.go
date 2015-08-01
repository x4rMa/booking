package database

import (
	"github.com/bborbe/booking/date"
	"github.com/jinzhu/gorm"
)

type Database interface {
	DB() (*gorm.DB, error)
}

type database struct {
	dbpath  string
	logmode bool
	db      *gorm.DB
}

func New(dbpath string, logmode bool) *database {
	s := new(database)
	s.dbpath = dbpath
	s.logmode = logmode
	s.db = nil
	return s
}

func (s *database) DB() (*gorm.DB, error) {
	if s.db == nil {
		db, err := gorm.Open("sqlite3", s.dbpath)
		if err != nil {
			return nil, err
		}
		db.SingularTable(true)
		db.LogMode(s.logmode)
		db.AutoMigrate(&date.Date{})
		s.db = &db
	}
	return s.db, nil
}
