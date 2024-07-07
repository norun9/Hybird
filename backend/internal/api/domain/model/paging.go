package model

type Paging struct {
	TotalPage  int `json:"totalPage"`
	Offset     int `json:"offset" validate:"gte=0"`
	Page       int `json:"page" validate:"gte=0"`
	Limit      int `json:"limit" validate:"gte=0"`
	TotalCount int `json:"totalCount"`
}
