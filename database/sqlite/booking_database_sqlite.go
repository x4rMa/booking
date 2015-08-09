package sqlite

import (
	booking_date "github.com/bborbe/booking/date"
	booking_model "github.com/bborbe/booking/model"
	booking_shooting "github.com/bborbe/booking/shooting"
	booking_user "github.com/bborbe/booking/user"
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
		db.AutoMigrate(&booking_user.User{}, &booking_date.Date{}, &booking_shooting.Shooting{}, &booking_model.Model{})
		s.db = &db
	}
	return s.db, nil
}
