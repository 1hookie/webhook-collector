package main

import (
	"encoding/json"
	"fmt"
	"time"

	"webhooks/model"
)

func main() {
	active := true

	h1 := model.Hook{ID: "abc", CreatedAt: time.Now()}
	h2 := model.Hook{ID: "def", CreatedAt: time.Now(), IsActive: &active}

	req := model.WebhookRequest{
		Method:     "POST",
		Path:       "/wh/abc",
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       `{"hello":"world"}`,
		IP:         "127.0.0.1",
		ReceivedAt: time.Now(),
	}

	for _, v := range []any{h1, h2, req} {
		b, _ := json.MarshalIndent(v, "", "  ")
		fmt.Println(string(b))
	}
}
