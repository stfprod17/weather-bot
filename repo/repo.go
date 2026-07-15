package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stfprod17/weatherbot/models"
)

type Repo struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}

}

// GetUserCity(ctx context.Context, userID int64) (string, error)
//
//	CreateUser(ctx context.Context, userID int64) error
//	UpdateCity(ctx context.Context, userID int64, city string) error
func (r *Repo) GetUserCity(ctx context.Context, userID int64) (string, error) {
	var city string

	row := r.db.QueryRow(ctx, "select coalesce(city, '') from users where ID=$1", userID)
	err := row.Scan(&city)
	if err != nil {
		return "", fmt.Errorf("error row.Scan: %w", err)
	}

	return city, nil
}

func (r *Repo) CreateUser(ctx context.Context, userID int64) error {
	_, err := r.db.Exec(ctx, "insert into users (id) values ($1)", userID)
	if err != nil {
		return fmt.Errorf("error db.Exec: %w ", err)
	}

	return nil
}

func (r *Repo) UpdateCity(ctx context.Context, userID int64, city string) error {
	_, err := r.db.Exec(ctx, "update users set city = $1 where id = $2", city, userID)
	if err != nil {
		return fmt.Errorf("error db.Exec: %w ", err)
	}

	return nil
}

func (r *Repo) GetUser(ctx context.Context, userID int64) (*models.User, error) {
	user := models.User{}

	row := r.db.QueryRow(ctx, "select id, coalesce(city, ''), coalesce(username, ''), created_at from users where id = $1", userID)
	err := row.Scan(&user.ID, &user.City, &user.Username, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error row.Scan: %w", err)
	}

	return &user, nil
}
