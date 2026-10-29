package config

import (
	"testing"
)

func TestConfig_validate(t *testing.T) {
	type fields struct {
		Port    int
		BaseURL string
		DB      DB
	}

	validDB := DB{
		Host:     "localhost",
		Port:     5432,
		DBName:   "test",
		User:     "test",
		Password: "test",
	}
	validURL := "http://localhost:8080"

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// --- порт приложения ---
		{"empty", fields{Port: 0, BaseURL: "", DB: DB{}}, true},
		{"negative port", fields{Port: -1, BaseURL: validURL, DB: validDB}, true},
		{"port too big", fields{Port: 65536, BaseURL: validURL, DB: validDB}, true},
		{"port min ok", fields{Port: 1, BaseURL: validURL, DB: validDB}, false},
		{"port max ok", fields{Port: 65535, BaseURL: validURL, DB: validDB}, false},

		// --- base_url ---
		{"empty base_url", fields{Port: 8080, BaseURL: "", DB: validDB}, true},
		{"base_url without scheme", fields{Port: 8080, BaseURL: "localhost:8080", DB: validDB}, true},
		{"base_url wrong scheme", fields{Port: 8080, BaseURL: "ftp://localhost", DB: validDB}, true},
		{"base_url without host", fields{Port: 8080, BaseURL: "http://", DB: validDB}, true},
		{"base_url https ok", fields{Port: 8080, BaseURL: "https://hooks.example.com", DB: validDB}, false},

		// --- секция db ---
		{"empty db", fields{Port: 8080, BaseURL: validURL, DB: DB{}}, true},
		{"empty db.port", fields{Port: 8080, BaseURL: validURL, DB: DB{
			Host:     "localhost",
			Port:     0,
			DBName:   "test",
			User:     "test",
			Password: "test",
		}}, true},
		{"db.port too big", fields{Port: 8080, BaseURL: validURL, DB: DB{
			Host:     "localhost",
			Port:     65536,
			DBName:   "test",
			User:     "test",
			Password: "test",
		}}, true},
		{"empty db.host", fields{Port: 8080, BaseURL: validURL, DB: DB{
			Host:     "",
			Port:     5432,
			DBName:   "test",
			User:     "test",
			Password: "test",
		}}, true},
		{"empty db.dbname", fields{Port: 8080, BaseURL: validURL, DB: DB{
			Host:     "localhost",
			Port:     5432,
			DBName:   "",
			User:     "test",
			Password: "test",
		}}, true},
		{"empty db.user", fields{Port: 8080, BaseURL: validURL, DB: DB{
			Host:     "localhost",
			Port:     5432,
			DBName:   "test",
			User:     "",
			Password: "test",
		}}, true},
		{"empty db.password", fields{Port: 8080, BaseURL: validURL, DB: DB{
			Host:     "localhost",
			Port:     5432,
			DBName:   "test",
			User:     "test",
			Password: "",
		}}, true},

		// --- всё корректно ---
		{"happy", fields{Port: 8080, BaseURL: validURL, DB: validDB}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Port:    tt.fields.Port,
				BaseURL: tt.fields.BaseURL,
				DB:      tt.fields.DB,
			}
			if err := c.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
