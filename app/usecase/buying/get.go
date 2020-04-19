package buying

import (
	"context"
	"time"

	"github.com/oharai/animal-crossings-turnip/app/domain/models/buyer"
)

func (u *UseCase) Get(ctx context.Context, count, page int64, now time.Time) ([]*buyer.Buyer, error) {
	return u.BuyerRepoFac.CreateRoRepository().SelectAvailable(ctx, count, page, now)
}
