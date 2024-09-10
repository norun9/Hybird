package transfer

import (
	"github.com/norun9/HyBird/backend/api/internal/domain/model"
	"github.com/norun9/HyBird/backend/api/lib/dbmodels"
)

func ToMessageEntity(m *model.Message) *dbmodels.Message {
	return (*dbmodels.Message)(m)
}
