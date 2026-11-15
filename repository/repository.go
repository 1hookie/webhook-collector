package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"webhooks/model"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CreateHook(ctx context.Context, hook model.Hook) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO hooks (id, created_at) VALUES ($1, $2)`,
		hook.ID, hook.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("create hook: %w", err)
	}
	return nil
}

func (r *Repo) HookExists(ctx context.Context, id string) (bool, error) {
	// Некорректный UUID тоже означает, что такого хука нет.
	if _, err := uuid.Parse(id); err != nil {
		return false, nil
	}

	var exists bool
	err := r.db.QueryRowContext(ctx,
		`SELECT EXISTS(SELECT 1 FROM hooks WHERE id = $1)`, id,
	).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check hook: %w", err)
	}
	return exists, nil
}
