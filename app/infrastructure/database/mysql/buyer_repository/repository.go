package buyer_repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/oharai/animal-crossings-turnip/app/domain/models/buyer"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) Insert(ctx context.Context, model *buyer.Buyer) (*buyer.Buyer, error) {
	stmt, err := r.db.Prepare("INSERT INTO Buyers (price, password, limit_num_visitor, created_at, expired_at) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(model.Price, model.Password, model.LimitNumVisitor, model.CreatedAt, model.ExpiredAt)
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

func (r *Repository) SelectAvailable(ctx context.Context, limit, page int64, now time.Time) ([]*buyer.Buyer, error) {
	stmt, err := r.db.Prepare(`
		SELECT * FROM Buyers
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

	var models []*buyer.Buyer
	for rows.Next() {
		model := buyer.Buyer{}
		err := rows.Scan(&model.ID, &model.Password, &model.Price, &model.LimitNumVisitor, &model.CreatedAt, &model.ExpiredAt)
		if err != nil {
			return nil, err
		}
		models = append(models, &model)
	}

	return models, nil
}
