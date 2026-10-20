package controller

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	mu    sync.Mutex
	hooks = map[string]time.Time{} // id хука -> время создания
)

// CreateHook — POST /api/v1/hooks создаёт хук и возвращает его id.
func CreateHook(c echo.Context) error {
	now := time.Now()
	id := strconv.FormatInt(now.UnixNano(), 10)

	mu.Lock()
	hooks[id] = now
	mu.Unlock()

	return c.String(http.StatusOK, id)
}

// CatchHook — ANY /wh/:hookId логирует запрос и отвечает 200.
func CatchHook(c echo.Context) error {
	r := c.Request()
	log.Println(r.Method, r.RequestURI, c.RealIP())

	return c.String(http.StatusOK, "OK")
}
