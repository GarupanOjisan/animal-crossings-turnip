package buyer

import "time"

type Buyer struct {
	ID              int64
	Price           int64
	Password        string
	LimitNumVisitor int64
	CreatedAt       *time.Time
	ExpiredAt       *time.Time
}
