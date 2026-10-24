package model

import "time"

type Hook struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"time"`
	IsActive  *bool     `json:"is_active,omitempty"`
}

type WebhookRequest struct {
	Method     string            `json:"method"`
	Path       string            `json:"path"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
	IP         string            `json:"ip"`
	ReceivedAt time.Time         `json:"received_at"`
}

func (h Hook) URL(baseURL string) string {
	return baseURL + "/wh/" + h.ID
}
