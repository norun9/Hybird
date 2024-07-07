package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/norun9/Hybird/pkg/dbmodels"
	"github.com/norun9/Hybird/pkg/util"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

type Query = qm.QueryMod

type Client struct {
	dbClient *sql.DB
}

func (c Client) Get(ctx context.Context) SQLHandler {
	if tx := util.GetDBTx(ctx); tx != nil {
		return tx
	}
	return c.dbClient
}

func NewDBClient(db *sql.DB) Client {
	initLocal()
	setBoil()
	return Client{db}
}

const location = "Asia/Tokyo"

func initLocal() {
	var loc *time.Location
	var err error
	if loc, err = time.LoadLocation(location); err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func setBoil() {
	boil.SetLocation(time.Local)
	boil.DebugMode = true
}

type SQLHandler interface {
	boil.ContextExecutor
}

func InnerJoin(outerTable, outerColumn, drivingTable, drivingColumn string) qm.QueryMod {
	return qm.InnerJoin(
		fmt.Sprintf(
			"%s on %s.%s = %s.%s",
			outerTable,
			drivingTable,
			drivingColumn,
			outerTable,
			outerColumn,
		))
}

func Count(ctx context.Context, dbClient SQLHandler, tableName string, queries []qm.QueryMod) (count int64, err error) {
	return CountByColumn(ctx, dbClient, tableName, "id", queries)
}

func CountByColumn(ctx context.Context, dbClient SQLHandler, tableName, column string, queries []qm.QueryMod) (count int64, err error) {
	totalCount := struct {
		Count int64
	}{}
	if err = dbmodels.NewQuery(append(queries,
		qm.Select(
			fmt.Sprintf(
				"count(distinct %s.%s) as count",
				tableName,
				column,
			)),
		qm.From(tableName),
	)...).Bind(ctx, dbClient, &totalCount); err != nil {
		return 0, err
	}
	return totalCount.Count, nil
}

func Distinct(tableName string) qm.QueryMod {
	return qm.Select(fmt.Sprintf("distinct %s.*", tableName))
}

func Select(tableName string, column string) qm.QueryMod {
	return qm.Select(fmt.Sprintf("%s.%s", tableName, column))
}

func OrderBy(column string, desc bool) qm.QueryMod {
	if desc {
		return qm.OrderBy(fmt.Sprintf("`%s` desc", column))
	}
	return qm.OrderBy(fmt.Sprintf("`%s`", column))
}

func GroupBy(tableName string, column string) qm.QueryMod {
	return qm.GroupBy(fmt.Sprintf("%s.%s", tableName, column))
}
