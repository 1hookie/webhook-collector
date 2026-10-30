package main

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"webhooks/config"
	"webhooks/controller"
	"webhooks/storage"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	store, err := storage.NewStore(100)
	if err != nil {
		log.Fatal(err)
	}

	h, err := controller.NewHandler(store, cfg.BaseURL)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.POST("/api/v1/hooks", h.CreateHook)
	e.GET("/api/v1/hooks/:id/requests", h.GetRequests)
	e.Any("/wh/:hookId", h.CatchHook)

	log.Printf("listening on :%d, base_url=%s", cfg.Port, cfg.BaseURL)
	log.Fatal(e.Start(":" + strconv.Itoa(cfg.Port)))
}
