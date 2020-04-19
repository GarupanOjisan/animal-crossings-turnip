package mysql

import (
	"database/sql"
	"fmt"
	"github.com/oharai/animal-crossings-turnip/app/domain/repository/seller_repository"
	mysql_seller_repo "github.com/oharai/animal-crossings-turnip/app/infrastructure/database/mysql/seller_repository"

	_ "github.com/go-sql-driver/mysql"

	"github.com/oharai/animal-crossings-turnip/app/domain/repository/buyer_repository"
	mysql_buyer_repo "github.com/oharai/animal-crossings-turnip/app/infrastructure/database/mysql/buyer_repository"
	"github.com/oharai/animal-crossings-turnip/config"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(c *config.Config) (*MySQL, error) {
	dbConf := c.Database.MySQL
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FTokyo",
		dbConf.User,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Database,
	))
	if err != nil {
		return nil, err
	}
	return &MySQL{db: db}, nil
}

// use only test
func (m *MySQL) GetDB() *sql.DB {
	return m.db
}

func (m *MySQL) GetBuyerRepoFactory() buyer_repository.Factory {
	return mysql_buyer_repo.NewFactory(m.db)
}

func (m *MySQL) GetSellerRepoFactory() seller_repository.Factory {
	return mysql_seller_repo.NewFactory(m.db)
}
