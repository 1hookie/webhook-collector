package cfgval

import "testing"

var validDB = DB{Host: "localhost", Port: 5432, DBName: "test", User: "test", Password: "test"}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{"happy", Config{Port: 8080, DB: validDB}, false},
		{"app port max ok", Config{Port: 65535, DB: validDB}, false},
		{"zero value", Config{}, true},
		{"app port zero", Config{Port: 0, DB: validDB}, true},
		{"app port too big", Config{Port: 65536, DB: validDB}, true},
		{"db host empty", Config{Port: 8080, DB: DB{Port: 5432, DBName: "test", User: "test", Password: "test"}}, true},
		{"db port zero", Config{Port: 8080, DB: DB{Host: "localhost", DBName: "test", User: "test", Password: "test"}}, true},
		{"db port too big", Config{Port: 8080, DB: DB{Host: "localhost", Port: 65536, DBName: "test", User: "test", Password: "test"}}, true},
		{"db name empty", Config{Port: 8080, DB: DB{Host: "localhost", Port: 5432, User: "test", Password: "test"}}, true},
		{"db user empty", Config{Port: 8080, DB: DB{Host: "localhost", Port: 5432, DBName: "test", Password: "test"}}, true},
		{"db password empty", Config{Port: 8080, DB: DB{Host: "localhost", Port: 5432, DBName: "test", User: "test"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if (err != nil) != tt.wantErr {
				t.Fatalf("Validate() = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
