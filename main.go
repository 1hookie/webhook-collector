package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"webhooks/controller"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/v1/hooks/:id/requests", controller.GetRequests)
	e.POST("/api/v1/hooks", controller.CreateHook)
	e.Any("/wh/:hookId", controller.CatchHook)

	log.Fatal(e.Start(":8080"))
}

// curl -i -X POST http://localhost:8080/api/v1/hooks
// curl -i http://localhost:8080/wh/1
// curl -i -X PUT "http://localhost:8080/wh/1?x=42"
