package selling

import (
	"context"
	"time"

	"github.com/oharai/animal-crossings-turnip/app/domain/models/seller"
)

func (u *UseCase) Get(ctx context.Context, count, page int64, now time.Time) ([]*seller.Seller, error) {
	return u.SellerRepoFac.CreateRoRepository().SelectAvailable(ctx, count, page, now)
}
