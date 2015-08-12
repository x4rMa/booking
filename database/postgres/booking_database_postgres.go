package postgres

import (
	"fmt"

	booking_date "github.com/bborbe/booking/date"
	booking_model "github.com/bborbe/booking/model"
	booking_shooting "github.com/bborbe/booking/shooting"
	booking_user "github.com/bborbe/booking/user"
	"github.com/bborbe/log"
	"github.com/jinzhu/gorm"
)

var logger = log.DefaultLogger

type postgres struct {
	dbname     string
	dbuser     string
	dbpassword string
	logmode    bool
	db         *gorm.DB
}

func New(dbname string, dbuser string, dbpassword string, logmode bool) *postgres {
	s := new(postgres)
	s.dbuser = dbuser
	s.dbpassword = dbpassword
	s.dbname = dbname
	s.logmode = logmode
	s.db = nil
	return s
}

func (s *postgres) DB() (*gorm.DB, error) {
	if s.db == nil {
		dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", s.dbuser, s.dbpassword, s.dbname)
		logger.Debugf("connect to postgres: %s", dbinfo)
		db, err := gorm.Open("postgres", dbinfo)
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
