package model

import (
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Paging
// https://github.com/volatiletech/sqlboiler/blob/master/queries/qm/query_mods.go
// Fulfill the QueryMod interface by implementing the Apply method
type Paging struct {
	TotalPage  int `json:"totalPage"`
	Offset     int `json:"offset" validate:"gte=0"`
	Page       int `json:"page" validate:"gte=0"`
	Limit      int `json:"limit" validate:"gte=0"`
	TotalCount int `json:"totalCount"`
}

func (p Paging) Apply(q *queries.Query) {
	if 0 < p.Limit {
		qm.Limit(p.Limit).Apply(q)
	}
	if 0 < p.Offset {
		qm.Offset(p.Offset).Apply(q)
	}
}
