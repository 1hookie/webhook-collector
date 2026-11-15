package envcfg

import "testing"

func TestFillFromEnv(t *testing.T) {
	base := App{Host: "app", Port: 8080}

	tests := []struct {
		name string
		env  map[string]string
		want App
	}{
		{"empty env", map[string]string{}, base},
		{"host", map[string]string{"HC_HOST": "api.local"}, App{Host: "api.local", Port: 8080}},
		{"port", map[string]string{"HC_PORT": "9000"}, App{Host: "app", Port: 9000}},
		{"both", map[string]string{"HC_HOST": "x", "HC_PORT": "1"}, App{Host: "x", Port: 1}},
		{"empty host ignored", map[string]string{"HC_HOST": ""}, base},
		{"bad port ignored", map[string]string{"HC_PORT": "abc"}, base},
		{"port zero ignored", map[string]string{"HC_PORT": "0"}, base},
		{"negative port ignored", map[string]string{"HC_PORT": "-1"}, base},
		{"port too big ignored", map[string]string{"HC_PORT": "65536"}, base},
		{"unknown vars ignored", map[string]string{"FOO": "1", "BAR": "2"}, base},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := base
			FillFromEnv(&cfg, tt.env)
			if cfg != tt.want {
				t.Fatalf("FillFromEnv(%v) = %+v, want %+v", tt.env, cfg, tt.want)
			}
		})
	}
}
