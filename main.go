package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"webhooks/controller"
)

func main() {
	e := echo.New()

	e.POST("/api/v1/hooks", controller.CreateHook)
	e.Any("/wh/:hookId", controller.CatchHook)

	log.Fatal(e.Start(":8080"))
}
