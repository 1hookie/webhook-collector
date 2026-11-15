package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"webhooks/config"
	"webhooks/conn"
	"webhooks/controller"
	"webhooks/repository"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	db, err := conn.NewConn(cfg.DB)
	if err != nil {
		return fmt.Errorf("connect to database: %w", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("close database: %v", err)
		}
	}()

	repo := repository.NewRepo(db)
	handler, err := controller.NewHandler(repo, cfg.BaseURL)
	if err != nil {
		return fmt.Errorf("create handler: %w", err)
	}

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.POST("/api/v1/hooks", handler.CreateHook)
	e.GET("/api/v1/hooks/:id/requests", handler.GetRequests)
	e.Any("/wh/:hookId", handler.CatchHook)

	log.Printf("listening on :%d, base_url=%s", cfg.Port, cfg.BaseURL)
	return e.Start(":" + strconv.Itoa(cfg.Port))
}
