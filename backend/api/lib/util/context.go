package util

import (
	"context"
	"database/sql"
)

type contextKey string

const dbTxContextKey contextKey = "dbTx"

func GetDBTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(dbTxContextKey).(*sql.Tx); ok {
		return tx
	}
	return nil
}
