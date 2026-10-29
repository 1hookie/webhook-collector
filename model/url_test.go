package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHookURL(t *testing.T) {
	const id = "6f1c2a9e-4b7d-4e2a-9c3f-1a2b3c4d5e6f"

	tests := []struct {
		name    string
		baseURL string
		id      string
		want    string
	}{
		{
			name:    "base without slash",
			baseURL: "http://localhost:8080",
			id:      id,
			want:    "http://localhost:8080/wh/" + id,
		},
		{
			name:    "base with trailing slash",
			baseURL: "http://localhost:8080/",
			id:      id,
			want:    "http://localhost:8080/wh/" + id,
		},
		{
			name:    "base with many trailing slashes",
			baseURL: "http://localhost:8080///",
			id:      id,
			want:    "http://localhost:8080/wh/" + id,
		},
		{
			name:    "id with leading slash",
			baseURL: "http://localhost:8080",
			id:      "/" + id,
			want:    "http://localhost:8080/wh/" + id,
		},
		{
			name:    "slashes on both sides",
			baseURL: "http://localhost:8080/",
			id:      "//" + id,
			want:    "http://localhost:8080/wh/" + id,
		},
		{
			name:    "https domain",
			baseURL: "https://hooks.example.com",
			id:      id,
			want:    "https://hooks.example.com/wh/" + id,
		},
		{
			name:    "base with path prefix",
			baseURL: "https://example.com/hookcatcher",
			id:      id,
			want:    "https://example.com/hookcatcher/wh/" + id,
		},
		{
			name:    "base with path prefix and trailing slash",
			baseURL: "https://example.com/hookcatcher/",
			id:      id,
			want:    "https://example.com/hookcatcher/wh/" + id,
		},
		{
			name:    "custom port from HC_PORT",
			baseURL: "http://localhost:8090",
			id:      id,
			want:    "http://localhost:8090/wh/" + id,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HookURL(tt.baseURL, tt.id)

			assert.Equal(t, tt.want, got)

			_, rest, _ := strings.Cut(got, "://")
			assert.NotContains(t, rest, "//", "в адресе не должно быть двойных слешей")
		})
	}
}

func TestHook_URL(t *testing.T) {
	h := Hook{ID: "abc"}
	assert.Equal(t, "http://localhost:8080/wh/abc", h.URL("http://localhost:8080/"))
}
