package main

import "testing"

func TestHeaderHas(t *testing.T) {
	tests := []struct {
		name  string
		head  string
		proto string
		want  bool
	}{
		{"happy", "application/json, text/html", "application/json", true},
		{"case-insensitive", "application/json, text/html", "Application/JSON", true},
		{"second part", "text/html, application/json", "application/json", true},
		{"with spaces", " text/html ,  application/json ", "application/json", true},
		{"not present", "text/html", "application/json", false},
		{"empty header", "", "application/json", false},
		{"only empty parts", " , ,", "application/json", false},
		{"unicode fold", "квай-гон, оби-ван", "КВАЙ-ГОН", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeaderHas(tt.head, tt.proto); got != tt.want {
				t.Fatalf("HeaderHas(%q, %q) = %v, want %v", tt.head, tt.proto, got, tt.want)
			}
		})
	}
}
