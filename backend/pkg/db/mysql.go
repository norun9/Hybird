package db

import (
	"database/sql"
	"fmt"
	"github.com/cockroachdb/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/norun9/Hybird/pkg/config"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const timezone string = "Asia%2FTokyo"

func NewDB(c config.DBConfig) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		c.User, c.Pass, c.Host, c.Port, c.Name, timezone)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect to database"))
	}
	boil.SetDB(db)
	return db
}
