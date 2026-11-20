package masker

import "testing"

func TestMaskEmail(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"email", "user@example.com", "****@***********"},
		{"short", "a@b", "*@*"},
		{"no at", "abc", "***"},
		{"cyrillic", "Оби-Ван", "*******"},
		{"empty", "", ""},
		{"double at", "a@@b", "*@@*"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaskEmail(tt.in); got != tt.want {
				t.Fatalf("MaskEmail(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
