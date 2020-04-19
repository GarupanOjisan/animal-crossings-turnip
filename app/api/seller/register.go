package seller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/oharai/animal-crossings-turnip/app/infrastructure/database/mysql"
	"github.com/oharai/animal-crossings-turnip/app/usecase/selling"
	"github.com/oharai/animal-crossings-turnip/config"
	"github.com/oharai/animal-crossings-turnip/constants"
)

type Endpoint struct {
	Conf *config.Config
}

var (
	ErrEmptyPrice       = fmt.Errorf("price must not be null")
	ErrInvalidPrice     = fmt.Errorf("invalid price")
	ErrInvalidPassword  = fmt.Errorf("invalid password")
	ErrInvalidLimit     = fmt.Errorf("invalid limit")
	ErrInvalidExpiredAt = fmt.Errorf("invalid expiredAt")
)

func (e *Endpoint) Register(c *gin.Context) {
	params, err := validateForRegister(c)
	if err != nil {
		log.Print(err)
		switch err {
		case ErrEmptyPrice:
		case ErrInvalidPrice:
		case ErrInvalidPassword:
		case ErrInvalidLimit:
		case ErrInvalidExpiredAt:
			c.Status(http.StatusUnprocessableEntity)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	db, err := mysql.NewMySQL(e.Conf)
	if err != nil {
		log.Print(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	u := selling.UseCase{
		SellerRepoFac: db.GetSellerRepoFactory(),
	}
	now := time.Now()
	b, err := u.Register(c, &selling.ValuesForRegister{
		Price:     params.Price,
		Password:  params.Password,
		ExpiredAt: params.ExpiredAt,
		CreatedAt: &now,
	})
	if err != nil {
		log.Print(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, b)
}

type validatedParamsForRegister struct {
	Price           int64
	LimitNumVisitor int64
	Password        string
	ExpiredAt       *time.Time
}

func validateForRegister(c *gin.Context) (*validatedParamsForRegister, error) {
	p := c.PostForm("price")
	if p == "" {
		return nil, ErrInvalidPrice
	}
	price, err := strconv.Atoi(p)
	if err != nil {
		return nil, err
	}
	if price < 0 {
		return nil, ErrInvalidPrice
	}

	l := c.PostForm("limit")
	var limit int
	if l != "" {
		limit, err = strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		if limit <= 0 {
			return nil, ErrInvalidLimit
		}
	}
	ex := c.PostForm("expired_at")
	var expiredAt *time.Time
	if ex != "" {
		t, err := time.Parse(time.RFC3339, ex)
		if err != nil {
			return nil, ErrInvalidExpiredAt
		}
		expiredAt = &t
	}

	password := c.PostForm("password")
	if password == "" || len(password) != constants.PASSWORD_LENGTH {
		return nil, ErrInvalidPassword
	}

	return &validatedParamsForRegister{
		Price:           int64(price),
		LimitNumVisitor: int64(limit),
		Password:        password,
		ExpiredAt:       expiredAt,
	}, nil
}
