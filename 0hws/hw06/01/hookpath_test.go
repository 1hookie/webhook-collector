package hookpath

import "testing"

func TestValidHookPath(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"hex id", "/wh/4f2c9a", true},
		{"numeric id", "/wh/123", true},
		{"mixed id with digit", "/wh/abc-7", true},
		{"empty id", "/wh/", false},
		{"id without digits", "/wh/abcd", false},
		{"wrong prefix", "/api/v1/hook", false},
		{"prefix without slash", "/wh123", false},
		{"empty string", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidHookPath(tt.in); got != tt.want {
				t.Fatalf("ValidHookPath(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
