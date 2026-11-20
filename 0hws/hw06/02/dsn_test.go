package dsn

import (
	"reflect"
	"testing"
)

func TestParseDSN(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want map[string]string
	}{
		{"happy", "host=localhost port=5432 user=test", map[string]string{"host": "localhost", "port": "5432", "user": "test"}},
		{"single pair", "sslmode=disable", map[string]string{"sslmode": "disable"}},
		{"empty value", "db= user=test", map[string]string{"db": "", "user": "test"}},
		{"token without equals", "garbage host=x", map[string]string{"host": "x"}},
		{"empty dsn", "", map[string]string{}},
		{"multiple spaces", "  a=1   b=2  ", map[string]string{"a": "1", "b": "2"}},
		{"value with equals", "a=b=c", map[string]string{"a": "b=c"}},
		{"empty key ignored", "=v x=1", map[string]string{"x": "1"}},
		{"last wins", "a=1 a=2", map[string]string{"a": "2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDSN(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("ParseDSN(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
