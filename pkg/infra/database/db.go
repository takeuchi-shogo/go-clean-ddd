package database

import (
	"errors"
	"fmt"

	"github.com/takeuchi-shogo/go-clean-ddd/pkg/infra/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MysqlDSN    = "%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true"
	PostgresDSN = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"

	ErrStrUnkownDBMS = "unknown DBMS: %s"
)

const (
	DBMSMySQL      = "mysql"
	DBMSPostgreSQL = "postgres"
)

func dsn(dbms string, c config.Database) (string, error) {
	switch dbms {
	case DBMSMySQL:
		return fmt.Sprintf(MysqlDSN, c.UserName, c.Password, c.Host, c.Port, c.DBName), nil
	case DBMSPostgreSQL:
		return fmt.Sprintf(PostgresDSN, c.UserName, c.Password, c.Host, c.Port, c.DBName), nil
	default:
		return "", errors.New(ErrStrUnkownDBMS)
	}
}

func dialector(dbms string, c config.Database) (gorm.Dialector, error) {
	dsn, err := dsn(dbms, c)
	if err != nil {
	}
	switch dbms {
	case DBMSMySQL:
		return mysql.Open(dsn), nil
	case DBMSPostgreSQL:
		return postgres.Open(dsn), nil
	default:
		return nil, errors.New(ErrStrUnkownDBMS)
	}
}

func NewDB(c config.DBConfig) (*gorm.DB, error) {
	dlctr, err := dialector(c.DBMS, c.Database)
	if err != nil {
		return nil, err
	}

	var db *gorm.DB
	db, err = gorm.Open(dlctr, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
	}

	sqlDB.SetMaxIdleConns(c.MaxIdleConn)
	sqlDB.SetMaxOpenConns(c.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(c.MaxLifetime)

	return db, nil
}
