package transfer

import (
	"github.com/norun9/Hybird/internal/api/domain/model"
	"github.com/norun9/Hybird/pkg/dbmodels"
)

func ToMessageEntity(m *model.Message) *dbmodels.Message {
	return (*dbmodels.Message)(m)
}
