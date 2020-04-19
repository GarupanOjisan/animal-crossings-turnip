package seller_repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/oharai/animal-crossings-turnip/app/domain/models/seller"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) Insert(ctx context.Context, model *seller.Seller) (*seller.Seller, error) {
	stmt, err := r.db.Prepare("INSERT INTO Sellers (price, password, created_at, expired_at) VALUES (?, ?, ?, ?);")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(model.Price, model.Password, model.CreatedAt, model.ExpiredAt)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("failed to insert row")
	}
	model.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (r *Repository) SelectAvailable(ctx context.Context, limit, page int64, now time.Time) ([]*seller.Seller, error) {
	stmt, err := r.db.Prepare(`
		SELECT * FROM Sellers
		WHERE
			expired_at is NULL OR
			expired_at >= ?
		ORDER BY created_at DESC, id
		LIMIT ?
		OFFSET ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(now.Format(time.RFC3339), limit, limit*page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models []*seller.Seller
	for rows.Next() {
		model := seller.Seller{}
		err := rows.Scan(&model.ID, &model.Password, &model.Price, &model.CreatedAt, &model.ExpiredAt)
		if err != nil {
			return nil, err
		}
		models = append(models, &model)
	}

	return models, nil
}
