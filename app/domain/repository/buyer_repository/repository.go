package buyer_repository

import (
	"context"
	"time"

	"github.com/oharai/animal-crossings-turnip/app/domain/models/buyer"
)

type Repository interface {
	Insert(ctx context.Context, model *buyer.Buyer) (*buyer.Buyer, error)
	SelectAvailable(ctx context.Context, limit, page int64, now time.Time) ([]*buyer.Buyer, error)
}
