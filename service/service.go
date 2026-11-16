package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"webhooks/entity"
	"webhooks/model"
	"webhooks/repository"
)

// Контроллер сможет проверять эту ошибку, не импортируя repository.
var ErrHookNotFound = repository.ErrHookNotFound

type Service struct {
	repo *repository.Repo
}

func NewService(repo *repository.Repo) (*Service, error) {
	if repo == nil {
		return nil, errors.New("repo is nil")
	}
	return &Service{repo: repo}, nil
}

func (s *Service) CreateHook(ctx context.Context, hook model.Hook) error {
	return s.repo.CreateHook(ctx, entity.Hook{
		ID:        hook.ID,
		CreatedAt: hook.CreatedAt,
	})
}

func (s *Service) HookExists(ctx context.Context, id string) (bool, error) {
	return s.repo.HookExists(ctx, id)
}

func (s *Service) AddRequest(
	ctx context.Context,
	hookID string,
	req model.WebhookRequest,
) error {
	headers, err := json.Marshal(req.Headers)
	if err != nil {
		return fmt.Errorf("encode headers: %w", err)
	}

	return s.repo.AddRequest(ctx, entity.WebhookRequest{
		HookID:     hookID,
		Method:     req.Method,
		Path:       req.Path,
		Headers:    json.RawMessage(headers),
		Body:       req.Body,
		IP:         req.IP,
		ReceivedAt: req.ReceivedAt,
	})
}

func (s *Service) ListRequests(
	ctx context.Context,
	hookID string,
) ([]model.WebhookRequest, error) {
	records, err := s.repo.ListRequests(ctx, hookID)
	if err != nil {
		return nil, err
	}

	result := make([]model.WebhookRequest, 0, len(records))
	for _, record := range records {
		var headers map[string]string
		if err := json.Unmarshal(record.Headers, &headers); err != nil {
			return nil, fmt.Errorf("decode headers for request %d: %w", record.ID, err)
		}

		result = append(result, model.WebhookRequest{
			Method:     record.Method,
			Path:       record.Path,
			Headers:    headers,
			Body:       record.Body,
			IP:         record.IP,
			ReceivedAt: record.ReceivedAt,
		})
	}
	return result, nil
}
