package seller_repository

import (
	"context"
	"time"

	"github.com/oharai/animal-crossings-turnip/app/domain/models/seller"
)

type Repository interface {
	Insert(ctx context.Context, model *seller.Seller) (*seller.Seller, error)
	SelectAvailable(ctx context.Context, limit, page int64, now time.Time) ([]*seller.Seller, error)
}
