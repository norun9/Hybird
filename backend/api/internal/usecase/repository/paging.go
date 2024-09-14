package repository

import "github.com/norun9/HyBird/backend/api/internal/domain/model"

// IPaging :
type IPaging interface {
	Get(totalCount, offset, page, limit int) (result model.Paging)
}
