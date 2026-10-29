package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type DB struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"DBName"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	Port    int    `json:"port"`
	BaseURL string `json:"base_url"`
	DB      DB     `json:"db"`
}

// NewConfig: файл → env → дефолты → валидация.
func NewConfig(fileName string) (*Config, error) {
	body, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("read config %s: %w", fileName, err)
	}

	var cfg Config
	if err := json.Unmarshal(body, &cfg); err != nil {
		return nil, fmt.Errorf("parse config %s: %w", fileName, err)
	}

	cfg.fillFromEnv()
	cfg.applyDefaults()

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}
	return &cfg, nil
}

// fillFromEnv переопределяет значения из файла переменными окружения.
// Некорректные значения молча игнорируются — остаётся то, что было в файле.
func (c *Config) fillFromEnv() {
	if v := os.Getenv("HC_PORT"); v != "" {
		if p, err := strconv.Atoi(v); err == nil && p > 0 {
			c.Port = p
		}
	}
	if v := os.Getenv("HC_BASE_URL"); v != "" {
		c.BaseURL = v
	}

	if v := os.Getenv("HC_DB_HOST"); v != "" {
		c.DB.Host = v
	}
	if v := os.Getenv("HC_DB_PORT"); v != "" {
		if p, err := strconv.Atoi(v); err == nil && p > 0 {
			c.DB.Port = p
		}
	}
	if v := os.Getenv("HC_DB_NAME"); v != "" {
		c.DB.DBName = v
	}
	if v := os.Getenv("HC_DB_USER"); v != "" {
		c.DB.User = v
	}
	if v := os.Getenv("HC_DB_PASSWORD"); v != "" {
		c.DB.Password = v
	}
}

// applyDefaults: если base_url не задан — собираем его из порта.
// Вызывается ПОСЛЕ fillFromEnv, чтобы учесть HC_PORT.
func (c *Config) applyDefaults() {
	if c.BaseURL == "" && c.Port > 0 {
		c.BaseURL = "http://localhost:" + strconv.Itoa(c.Port)
	}
}

func (c *Config) validate() error {
	if c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("port must be in 1..65535, got %d", c.Port)
	}

	u, err := url.Parse(c.BaseURL)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return fmt.Errorf("base_url must be an absolute http(s) URL, got %q", c.BaseURL)
	}

	if c.DB.Host == "" {
		return fmt.Errorf("db.host must be specified")
	}
	if c.DB.Port < 1 || c.DB.Port > 65535 {
		return fmt.Errorf("db.port must be in 1..65535, got %d", c.DB.Port)
	}
	if c.DB.DBName == "" {
		return fmt.Errorf("db.DBName must be specified")
	}
	if c.DB.User == "" {
		return fmt.Errorf("db.user must be specified")
	}
	if c.DB.Password == "" {
		return fmt.Errorf("db.password must be specified")
	}
	return nil
}
