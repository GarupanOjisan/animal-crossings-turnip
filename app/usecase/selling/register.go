package selling

import (
	"context"
	"time"

	"github.com/oharai/animal-crossings-turnip/app/domain/models/seller"
	"github.com/oharai/animal-crossings-turnip/app/domain/repository/seller_repository"
)

type UseCase struct {
	SellerRepoFac seller_repository.Factory
}

type ValuesForRegister struct {
	ID        int64
	Price     int64
	Password  string
	CreatedAt *time.Time
	ExpiredAt *time.Time
}

func (u *UseCase) Register(ctx context.Context, v *ValuesForRegister) (*seller.Seller, error) {
	return u.SellerRepoFac.CreateRwRepository().Insert(ctx, &seller.Seller{
		Password:  v.Password,
		Price:     v.Price,
		CreatedAt: v.CreatedAt,
		ExpiredAt: v.ExpiredAt,
	})
}
