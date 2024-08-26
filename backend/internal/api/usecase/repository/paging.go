package repository

import "github.com/norun9/Hybird/internal/api/domain/model"

// IPaging :
type IPaging interface {
	Get(totalCount, offset, page, limit int) (result model.Paging)
}
