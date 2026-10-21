package model

import "time"

type Hook struct {
	Id         int       `json:"id"`
	Created_at time.Time `json:"time"`
	IsActive   *bool     `json:"status"`
}

type WebHookRequest struct {
	Method      string            `json:"method"`
	Path        string            `json:"path"`
	Headers     map[string]string `json:"headers"`
	Body        string            `json:"body"`
	Ip          string            `json:"ip"`
	Recieved_at time.Time         `json:"recieved_at"`
}
