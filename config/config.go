package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

const defaultConfigFile = "config.json"

// имена переменных окружения.
const (
	envConfigFile = "HC_CONFIG_FILE"
	envPort       = "HC_PORT"
	envBaseURL    = "HC_BASE_URL"
	envDBHost     = "HC_DB_HOST"
	envDBPort     = "HC_DB_PORT"
	envDBName     = "HC_DB_NAME"
	envDBUser     = "HC_DB_USER"
	envDBPassword = "HC_DB_PASSWORD"
)

// envKeys переменные, которые переопределяют значения из файла.
var envKeys = []string{
	envPort, envBaseURL,
	envDBHost, envDBPort,
	envDBName, envDBUser,
	envDBPassword,
}

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

// options  параметры загрузки конфига
type options struct {
	file   string
	getenv func(string) string
}

type Option func(*options)

// WithFile задаёт путь к файлу конфигурации.
func WithFile(path string) Option {
	return func(o *options) {
		o.file = path
	}
}

// WithEnv задаёт источник переменных окружения
func WithEnv(getenv func(string) string) Option {
	return func(o *options) {
		o.getenv = getenv
	}
}

// NewConfig: файл -> env -> дефолты -> валидация.
func NewConfig(opts ...Option) (*Config, error) {
	o := options{getenv: os.Getenv}
	for _, opt := range opts {
		opt(&o)
	}

	// путь к файлу: опция -> HC_CONFIG_FILE -> дефолт
	file := o.file
	switch {
	case file != "":
		// задан опцией WithFile — оставляем
	case o.getenv(envConfigFile) != "":
		file = o.getenv(envConfigFile)
	default:
		file = defaultConfigFile
	}

	body, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("read config %s: %w", file, err)
	}

	var cfg Config
	if err := json.Unmarshal(body, &cfg); err != nil {
		return nil, fmt.Errorf("parse config %s: %w", file, err)
	}

	cfg.fillFromEnv(o.getenv)
	cfg.applyDefaults()

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}
	return &cfg, nil
}

// fillFromEnv переопределяет значения из файла переменными окружения.
func (c *Config) fillFromEnv(getenv func(string) string) {
	for _, key := range envKeys {
		v := getenv(key)
		if v == "" {
			continue
		}

		switch key {
		case envPort:
			if p, ok := parsePort(v); ok {
				c.Port = p
			}
		case envBaseURL:
			c.BaseURL = v
		case envDBHost:
			c.DB.Host = v
		case envDBPort:
			if p, ok := parsePort(v); ok {
				c.DB.Port = p
			}
		case envDBName:
			c.DB.DBName = v
		case envDBUser:
			c.DB.User = v
		case envDBPassword:
			c.DB.Password = v
		}
	}
}

// parsePort возвращает порт и true, если строка — целое число в 1..65535.
func parsePort(s string) (int, bool) {
	p, err := strconv.Atoi(s)
	switch {
	case err != nil:
		return 0, false
	case p < 1 || p > 65535:
		return 0, false
	default:
		return p, true
	}
}

func (c *Config) applyDefaults() {
	if c.BaseURL == "" && c.Port > 0 {
		c.BaseURL = "http://localhost:" + strconv.Itoa(c.Port)
	}
}

func (c *Config) validate() error {
	switch {
	case c.Port < 1 || c.Port > 65535:
		return fmt.Errorf("port must be in 1..65535, got %d", c.Port)
	case !isHTTPURL(c.BaseURL):
		return fmt.Errorf("base_url must be an absolute http(s) URL, got %q", c.BaseURL)
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

// абсолютный URL со схемой http/https и непустым хостом.
func isHTTPURL(s string) bool {
	u, err := url.Parse(s)
	if err != nil || u.Host == "" {
		return false
	}
	switch u.Scheme {
	case "http", "https":
		return true
	default:
		return false
	}
}
