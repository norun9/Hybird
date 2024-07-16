package db

import (
	"database/sql"
	"fmt"
	"github.com/cockroachdb/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/norun9/Hybird/pkg/config"
	"github.com/norun9/Hybird/pkg/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"
	_ "time/tzdata"
)

const timezone string = "Asia%2FTokyo"

func NewDB(c config.DBConfig) *sql.DB {
	log.Logger.Info(fmt.Sprintf("Config:%v", c))
	loc, err := time.LoadLocation(timezone)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		c.User, c.Pass, c.Host, c.Name, loc)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect to database"))
	}

	if err = db.Ping(); err != nil {
		panic(errors.Wrap(err, "failed to connect to database"))
	}

	boil.SetDB(db)

	return db
}
