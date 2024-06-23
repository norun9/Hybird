package main

import (
	"fmt"
	"github.com/norun9/Hybird/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const timezone string = "Asia%2FTokyo"

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./pkg/db/dao",
		Mode:              gen.WithDefaultQuery, // generate mode
		ModelPkgPath:      "dao",
		WithUnitTest:      true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		FieldNullable:     true,
	})
	c := config.Prepare()
	dbConf := c.DBConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.Name, timezone)
	db, _ := gorm.Open(mysql.Open(dsn))
	g.UseDB(db)
	// NOTE: Please remember to write `GenerateModel(tableName)` code whenever adding new tables to the database.
	g.GenerateModel("messages")
	g.Execute()
}
