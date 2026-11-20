package substr

import "testing"

func TestSubstringStats(t *testing.T) {
	tests := []struct {
		name string
		s    string
		sub  string
		want Stats
	}{
		{"happy", "hello world", "world", Stats{1, 6, 6}},
		{"non-overlapping", "aaaa", "aa", Stats{2, 0, 2}},
		{"repeat", "абв абв абв", "абв", Stats{3, 0, 14}},
		{"cyrillic bytes", "Падме Амидала", "Амидала", Stats{1, 11, 11}},
		{"not found", "abc", "z", Stats{0, -1, -1}},
		{"empty s", "", "a", Stats{0, -1, -1}},
		{"empty sub", "abc", "", Stats{4, 0, 3}},
		{"whole string", "Падме", "Падме", Stats{1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubstringStats(tt.s, tt.sub); got != tt.want {
				t.Fatalf("SubstringStats(%q, %q) = %+v, want %+v", tt.s, tt.sub, got, tt.want)
			}
		})
	}
}
