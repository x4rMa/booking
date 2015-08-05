package sqlite

import (
	"github.com/jinzhu/gorm"
)

type sqlite struct {
	dbpath  string
	logmode bool
	db      *gorm.DB
}

func New(dbpath string, logmode bool) *sqlite {
	s := new(sqlite)
	s.dbpath = dbpath
	s.logmode = logmode
	s.db = nil
	return s
}

func (s *sqlite) DB() (*gorm.DB, error) {
	if s.db == nil {
		db, err := gorm.Open("sqlite3", s.dbpath)
		if err != nil {
			return nil, err
		}
		db.SingularTable(true)
		db.LogMode(s.logmode)
		s.db = &db
	}
	return s.db, nil
}