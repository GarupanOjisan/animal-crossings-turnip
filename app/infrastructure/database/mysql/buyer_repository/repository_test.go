package buyer_repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/oharai/animal-crossings-turnip/app/infrastructure/database/mysql/buyer_repository"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/oharai/animal-crossings-turnip/app/domain/models/buyer"
	"github.com/oharai/animal-crossings-turnip/app/infrastructure/database/mysql"
	"github.com/oharai/animal-crossings-turnip/config"
)

func TestRepository_Insert(t *testing.T) {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := mysql.NewMySQL(conf)
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx   context.Context
		model *buyer.Buyer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *buyer.Buyer
		wantErr bool
	}{
		{
			name:   "success inserting a single row",
			fields: fields{db: db.GetDB()},
			args: args{
				ctx: nil,
				model: &buyer.Buyer{
					Password:        "test1",
					Price:           100,
					LimitNumVisitor: -1,
					CreatedAt:       &now,
					ExpiredAt:       nil,
				},
			},
			want: &buyer.Buyer{
				Password:        "test1",
				Price:           100,
				LimitNumVisitor: -1,
				CreatedAt:       &now,
				ExpiredAt:       nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fac := db.GetBuyerRepoFactory()
			r := fac.CreateRwRepository()
			got, err := r.Insert(tt.args.ctx, tt.args.model)
			defer clean(t, tt.fields.db, []*buyer.Buyer{got})
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(buyer.Buyer{}, "ID"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("SelectAvailable() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func clean(t *testing.T, db *sql.DB, buyers []*buyer.Buyer) {
	ids := make([]interface{}, len(buyers))
	for i, b := range buyers {
		ids[i] = b.ID
	}

	params := ""
	for _, _ = range buyers {
		params += "?,"
	}
	stmt, err := db.Prepare(fmt.Sprintf("DELETE FROM Buyers WHERE id IN (%s)", strings.TrimSuffix(params, ",")))
	if err != nil {
		t.Fatal(err)
	}
	_, err = stmt.Exec(ids...)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRepository_SelectAvailable(t *testing.T) {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := mysql.NewMySQL(conf)
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()

	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx   context.Context
		limit int64
		page  int64
		now   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    []*buyer.Buyer
		want    []*buyer.Buyer
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				db: db.GetDB(),
			},
			args: args{
				ctx:   nil,
				limit: 2,
				page:  1,
				now:   now,
			},
			mock: []*buyer.Buyer{
				{
					Price:           101,
					Password:        "10000",
					LimitNumVisitor: -1,
					CreatedAt:       &now,
					ExpiredAt:       nil,
				},
				{
					Price:           102,
					Password:        "10001",
					LimitNumVisitor: 10,
					CreatedAt:       &now,
					ExpiredAt:       nil,
				},
				{
					Price:           103,
					Password:        "10001",
					LimitNumVisitor: 10,
					CreatedAt:       &now,
					ExpiredAt: func() *time.Time {
						t := now.Add(-time.Hour)
						return &t
					}(),
				},
				{
					Price:           104,
					Password:        "10000",
					LimitNumVisitor: -1,
					CreatedAt:       &now,
					ExpiredAt:       nil,
				},
				{
					Price:           105,
					Password:        "10001",
					LimitNumVisitor: 10,
					CreatedAt:       &now,
					ExpiredAt: func() *time.Time {
						t := now.Add(-time.Hour)
						return &t
					}(),
				},
				{
					Price:           106,
					Password:        "10000",
					LimitNumVisitor: -1,
					CreatedAt:       &now,
					ExpiredAt: func() *time.Time {
						t := now.Add(time.Hour)
						return &t
					}(),
				},
			},
			want: []*buyer.Buyer{
				{
					Price:           104,
					Password:        "10000",
					LimitNumVisitor: -1,
					CreatedAt:       &now,
					ExpiredAt:       nil,
				},
				{
					Price:           106,
					Password:        "10000",
					LimitNumVisitor: -1,
					CreatedAt:       &now,
					ExpiredAt: func() *time.Time {
						t := now.Add(time.Hour)
						return &t
					}(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer setup(t, tt.fields.db, tt.mock)()
			r := buyer_repository.NewFactory(tt.fields.db).CreateRwRepository()
			got, err := r.SelectAvailable(tt.args.ctx, tt.args.limit, tt.args.page, tt.args.now)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectAvailable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(buyer.Buyer{}, "ID"),
				cmpopts.IgnoreFields(buyer.Buyer{}, "CreatedAt"),
				cmpopts.IgnoreFields(buyer.Buyer{}, "ExpiredAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("SelectAvailable() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func setup(t *testing.T, db *sql.DB, mock []*buyer.Buyer) func() {
	sql := "INSERT INTO Buyers (price, password, limit_num_visitor, created_at, expired_at) VALUES "

	var values []interface{}
	for _, m := range mock {
		sql += "(?, ?, ?, ?, ?),"
		values = append(values, m.Price, m.Password, m.LimitNumVisitor, m.CreatedAt, m.ExpiredAt)
	}
	sql = strings.TrimSuffix(sql, ",")
	stmt, err := db.Prepare(sql)
	if err != nil {
		t.Fatal(err)
	}

	res, err := stmt.Exec(values...)
	if err != nil {
		t.Fatal(err)
	}

	return func() {
		for i, m := range mock {
			id, err := res.LastInsertId()
			if err != nil {
				t.Fatal(err)
			}
			m.ID = id + int64(i)
		}
		clean(t, db, mock)
	}
}
