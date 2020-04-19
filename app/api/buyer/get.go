package buyer

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/oharai/animal-crossings-turnip/app/infrastructure/database/mysql"
	"github.com/oharai/animal-crossings-turnip/app/usecase/buying"
)

var (
	ErrInvalidPage = fmt.Errorf("invalid page")
)

func (e *Endpoint) Get(c *gin.Context) {
	params, err := validateForGet(c)
	if err != nil {
		log.Print(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	db, err := mysql.NewMySQL(e.Conf)
	if err != nil {
		log.Print(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	u := buying.UseCase{BuyerRepoFac: db.GetBuyerRepoFactory()}
	buyers, err := u.Get(c, e.Conf.App.PageSize, params.Page, time.Now())
	if err != nil {
		log.Print(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, buyers)
}

type validatedParamsForGet struct {
	Page int64
}

func validateForGet(c *gin.Context) (*validatedParamsForGet, error) {
	p := c.Query("page")
	page := 0
	if p != "" {
		var err error
		page, err = strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		if page < 0 {
			return nil, ErrInvalidPage
		}
	}
	return &validatedParamsForGet{Page: int64(page)}, nil
}
