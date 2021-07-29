package db

import (
	"ImyouboHomeKit/config"
	"ImyouboHomeKit/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var EmptyResultErr = errors.DaoEmptyResultError()

type DB struct {
	*xorm.Engine
}

func GetDB(driver, url string) (*DB, error) {
	engine, err := xorm.NewEngine(driver, url)
	if err != nil {
		return nil, err
	}
	return &DB{engine}, nil
}

func GetDefaultDB() (*DB, error) {
	return GetDB(config.DataSourceConfig.Driver, config.DataSourceConfig.Url)
}



