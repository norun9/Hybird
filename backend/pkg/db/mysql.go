package db

import (
	"fmt"
	"github.com/norun9/Hybird/pkg/config"

	"github.com/cockroachdb/errors"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const timezone string = "Asia%2FTokyo"

func NewDB() *gorm.DB {
	c := config.Prepare()
	dbConf := c.DBConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.Name, timezone)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // data source name
	}), &gorm.Config{})

	if err != nil {
		panic(errors.Wrap(err, "failed to connect to database"))
	}

	return db
}
