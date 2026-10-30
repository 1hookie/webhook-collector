// Задача 4. «Валидация конфига» — многоправильная проверка вложенной структуры.
package cfgval

import "fmt"

type DB struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

type Config struct {
	Port int
	DB   DB
}

// Validate проверяет поля по одному: одна проверка — одно сообщение.
// Value-ресивер: метод ничего не меняет, а Config — маленькая структура без мьютекса.
func (c Config) Validate() error {
	switch {
	case c.Port < 1 || c.Port > 65535:
		return fmt.Errorf("port must be in 1..65535, got %d", c.Port)
	case c.DB.Host == "":
		return fmt.Errorf("db.host must be specified")
	case c.DB.Port < 1 || c.DB.Port > 65535:
		return fmt.Errorf("db.port must be in 1..65535, got %d", c.DB.Port)
	case c.DB.DBName == "":
		return fmt.Errorf("db.DBName must be specified")
	case c.DB.User == "":
		return fmt.Errorf("db.user must be specified")
	case c.DB.Password == "":
		return fmt.Errorf("db.password must be specified")
	default:
		return nil
	}
}
