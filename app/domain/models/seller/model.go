package seller

import "time"

type Seller struct {
	ID        int64
	Price     int64
	Password  string
	CreatedAt *time.Time
	ExpiredAt *time.Time
}
