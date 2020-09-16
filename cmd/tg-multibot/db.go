package main

import (
	"errors"
	"fmt"

	// log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// initDatabase function for initialize database
func InitDatabase() (db *gorm.DB, err error) {
	if options.DBDriver == "mysql" {
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", options.DBUser, options.DBPassword, options.DBHost, options.DBPort, options.DBName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		// db = db.Set("gorm:table_options", "CHARSET=utf8")
	} else if options.DBDriver == "postgres" {
		dsn := fmt.Sprintf("user=%v password=%v DB.name=%v port=%v sslmode=disable TimeZone=Europe/London", options.DBUser, options.DBPassword, options.DBName, options.DBPort)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else if options.DBDriver == "sqlite" || options.DBDriver == "sqlite3" {
		db, err = gorm.Open(sqlite.Open(options.DBName), &gorm.Config{})
	} else {
		return nil, errors.New("not supported database adapter")
	}
	return
}
