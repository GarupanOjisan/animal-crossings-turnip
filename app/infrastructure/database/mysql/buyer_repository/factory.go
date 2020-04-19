package buyer_repository

import (
	"database/sql"

	"github.com/oharai/animal-crossings-turnip/app/domain/repository/buyer_repository"
)

type Factory struct {
	db *sql.DB
}

func NewFactory(db *sql.DB) *Factory {
	return &Factory{db: db}
}

func (f Factory) CreateRoRepository() buyer_repository.Repository {
	return &Repository{db: f.db}
}

func (f Factory) CreateRwRepository() buyer_repository.Repository {
	return &Repository{db: f.db}
}
