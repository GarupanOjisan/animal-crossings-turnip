package buying

import (
	"context"
	"time"

	"github.com/oharai/animal-crossings-turnip/app/domain/models/buyer"
	"github.com/oharai/animal-crossings-turnip/app/domain/repository/buyer_repository"
)

type UseCase struct {
	BuyerRepoFac buyer_repository.Factory
}

type ValuesForRegister struct {
	Price     int64
	Password  string
	Limit     int64
	CreatedAt *time.Time
	ExpiredAt *time.Time
}

func (u *UseCase) Register(ctx context.Context, v *ValuesForRegister) (*buyer.Buyer, error) {
	return u.BuyerRepoFac.CreateRwRepository().Insert(ctx, &buyer.Buyer{
		Password:        v.Password,
		Price:           v.Price,
		LimitNumVisitor: v.Limit,
		CreatedAt:       v.CreatedAt,
		ExpiredAt:       v.ExpiredAt,
	})
}
