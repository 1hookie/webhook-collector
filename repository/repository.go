package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"webhooks/model"
)

var ErrHookNotFound = errors.New("hook not found")

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

func (r *Repo) AddRequest(
	ctx context.Context,
	hookID string,
	req model.WebhookRequest,
) error {
	headers, err := json.Marshal(req.Headers)
	if err != nil {
		return fmt.Errorf("encode headers: %w", err)
	}

	_, err = r.db.ExecContext(ctx,
		`INSERT INTO requests (hook_id, method, path, headers, body, ip, received_at)
	VALUES ($1, $2, $3, $4::jsonb, $5, $6, $7)`, hookID, req.Method, req.Path, string(headers), req.Body, req.IP, req.ReceivedAt)
	if err != nil {
		return fmt.Errorf("save request: %w", err)
	}
	return nil
}

func (r *Repo) Requests(
	ctx context.Context,
	hookID string,
) (result []model.WebhookRequest, err error) {
	exists, err := r.HookExists(ctx, hookID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrHookNotFound
	}

	rows, err := r.db.QueryContext(ctx, `
		SELECT method, path, headers, body, ip, received_at
		FROM requests
		WHERE hook_id = $1
		ORDER BY id ASC
	`, hookID)
	if err != nil {
		return nil, fmt.Errorf("get requests: %w", err)
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close request rows: %w", closeErr)
		}
	}()

	result = make([]model.WebhookRequest, 0)
	for rows.Next() {
		var req model.WebhookRequest
		var headers []byte

		if scanErr := rows.Scan(
			&req.Method, &req.Path, &headers,
			&req.Body, &req.IP, &req.ReceivedAt,
		); scanErr != nil {
			return nil, fmt.Errorf("scan request: %w", scanErr)
		}
		if jsonErr := json.Unmarshal(headers, &req.Headers); jsonErr != nil {
			return nil, fmt.Errorf("decode headers: %w", jsonErr)
		}

		result = append(result, req)
	}
	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, fmt.Errorf("iterate requests: %w", rowsErr)
	}
	return result, nil
}
