package options

import "testing"

func TestNewServer(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		want Server
	}{
		{"defaults", nil, Server{Host: "localhost", Port: 8080, Timeout: 30, Workers: 4}},
		{"host", []Option{WithHost("api.local")}, Server{Host: "api.local", Port: 8080, Timeout: 30, Workers: 4}},
		{"port", []Option{WithPort(9000)}, Server{Host: "localhost", Port: 9000, Timeout: 30, Workers: 4}},
		{"timeout", []Option{WithTimeout(15)}, Server{Host: "localhost", Port: 8080, Timeout: 15, Workers: 4}},
		{"workers", []Option{WithWorkers(8)}, Server{Host: "localhost", Port: 8080, Timeout: 30, Workers: 8}},
		{"all", []Option{WithHost("x"), WithPort(9000), WithTimeout(15), WithWorkers(8)},
			Server{Host: "x", Port: 9000, Timeout: 15, Workers: 8}},
		{"last wins", []Option{WithPort(1), WithPort(2)}, Server{Host: "localhost", Port: 2, Timeout: 30, Workers: 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.opts...); got != tt.want {
				t.Fatalf("NewServer(%v) = %+v, want %+v", tt.opts, got, tt.want)
			}
		})
	}
}

func TestServerString(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		want string
	}{
		{"default", nil, "localhost:8080 t=30 w=4"},
		{"custom", []Option{WithHost("api.local"), WithPort(9000), WithTimeout(15), WithWorkers(8)},
			"api.local:9000 t=15 w=8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.opts...).String(); got != tt.want {
				t.Fatalf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}
