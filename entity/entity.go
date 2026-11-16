package entity

import (
	"encoding/json"
	"time"
)

type Hook struct {
	ID        string
	CreatedAt time.Time
}

type WebhookRequest struct {
	ID         int64
	HookID     string
	Method     string
	Path       string
	Headers    json.RawMessage
	Body       string
	IP         string
	ReceivedAt time.Time
}
