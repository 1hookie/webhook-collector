package conn

import (
	"database/sql"
	"fmt"
	"net"
	"net/url"
	"strconv"

	_ "github.com/lib/pq"

	"webhooks/config"
)

func NewConn(cfg config.DB) (*sql.DB, error) {
	dsn := (&url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Path:     "/" + cfg.DBName,
		RawQuery: "sslmode=disable&connect_timeout=5",
	}).String()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}
	return db, nil
}
