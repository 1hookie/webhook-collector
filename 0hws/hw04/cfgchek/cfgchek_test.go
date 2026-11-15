package cfgcheck

import "testing"

func TestParseFlags(t *testing.T) {
	def := "./config.json"

	tests := []struct {
		name    string
		args    []string
		want    Options
		wantErr bool
	}{
		{"defaults", nil, Options{ConfigFile: def}, false},
		{"config value", []string{"--config", "app.yml"}, Options{ConfigFile: "app.yml"}, false},
		{"config equals", []string{"--config=app.yml"}, Options{ConfigFile: "app.yml"}, false},
		{"pretty", []string{"--pretty"}, Options{ConfigFile: def, Pretty: true}, false},
		{"quiet", []string{"--quiet"}, Options{ConfigFile: def, Quiet: true}, false},
		{"all", []string{"--pretty", "--quiet", "--config", "prod.json"},
			Options{ConfigFile: "prod.json", Pretty: true, Quiet: true}, false},
		{"unknown flag", []string{"--nope"}, Options{}, true},
		{"positional arg", []string{"extra.json"}, Options{}, true},
		{"config without value", []string{"--config"}, Options{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFlags(tt.args)
			if (err != nil) != tt.wantErr {
				t.Fatalf("parseFlags(%v): err = %v, wantErr = %v", tt.args, err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Fatalf("parseFlags(%v) = %+v, want %+v", tt.args, got, tt.want)
			}
		})
	}
}
