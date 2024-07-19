package repository

import (
	"github.com/norun9/Hybird/internal/api/domain/model"
	"github.com/norun9/Hybird/internal/api/usecase/repository"
	"math"
)

var (
	defaultPageLimit = 20
)

type paging struct{}

func NewPaging() repository.Paging {
	return &paging{}
}

func (r *paging) Get(totalCount, offset, page, limit int) (output model.Paging) {
	if limit == 0 {
		limit = defaultPageLimit
	}
	totalPage := int(math.Ceil(float64(totalCount) / float64(limit)))
	// When pages are larger than total pages, set to the number of total pages.
	if totalPage < page {
		page = totalPage
	}
	if page == 0 {
		page = 1
	}
	offset = limit*(page-1) + offset

	return model.Paging{
		TotalPage:  totalPage,
		Offset:     offset,
		Page:       page,
		Limit:      limit,
		TotalCount: totalCount,
	}
}
