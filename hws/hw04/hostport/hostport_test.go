package hostport

import "testing"

func TestParseHostPort(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		wantHost string
		wantPort int
		wantErr  bool
	}{
		{"happy", "localhost:8080", "localhost", 8080, false},
		{"ip host", "192.168.1.1:22", "192.168.1.1", 22, false},
		{"min port", "db:1", "db", 1, false},
		{"max port", "db:65535", "db", 65535, false},
		{"empty string", "", "", 0, true},
		{"no colon", "localhost", "", 0, true},
		{"empty host", ":8080", "", 0, true},
		{"empty port", "localhost:", "", 0, true},
		{"port zero", "localhost:0", "", 0, true},
		{"port too big", "localhost:65536", "", 0, true},
		{"negative port", "localhost:-5", "", 0, true},
		{"port not a number", "localhost:8o8", "", 0, true},
		{"two colons", "a:b:80", "", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			host, port, err := ParseHostPort(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseHostPort(%q): err = %v, wantErr = %v", tt.in, err, tt.wantErr)
			}
			if !tt.wantErr && (host != tt.wantHost || port != tt.wantPort) {
				t.Fatalf("ParseHostPort(%q) = (%q, %d), want (%q, %d)", tt.in, host, port, tt.wantHost, tt.wantPort)
			}
		})
	}
}
