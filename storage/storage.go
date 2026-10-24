package storage

import (
	"errors"
	"sync"
	"webhooks/model"
)

var ErrHookNotFound = errors.New("hook not found")

type Store struct {
	mu         sync.RWMutex
	hooks      map[string]model.Hook             // id хука -> хук
	requests   map[string][]model.WebhookRequest // id хука -> его запросы
	maxPerHook int
}

func NewStore(maxPerHook int) (*Store, error) {
	if maxPerHook <= 0 {
		return nil, errors.New("maxPerHook must be greater than 0")
	}

	return &Store{
		hooks:      make(map[string]model.Hook),
		requests:   make(map[string][]model.WebhookRequest),
		maxPerHook: maxPerHook,
	}, nil
}

func (s *Store) CreateHook(h model.Hook) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.hooks[h.ID] = h
	s.requests[h.ID] = []model.WebhookRequest{}
}

func (s *Store) AddRequst(hookID string, req model.WebhookRequest) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.hooks[hookID]; !ok {
		return ErrHookNotFound
	}
	reqs := s.requests[hookID]
	if len(reqs) >= s.maxPerHook {
		reqs = reqs[1:]
	}

	s.requests[hookID] = append(reqs, req)
	return nil
}

func (s *Store) Requests(hookID string) ([]model.WebhookRequest, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, ok := s.hooks[hookID]; !ok {
		return nil, ErrHookNotFound
	}

	src := s.requests[hookID]
	out := make([]model.WebhookRequest, len(src))
	copy(out, src)
	return out, nil
}
