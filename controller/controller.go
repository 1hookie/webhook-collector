package controller

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"webhooks/model"
)

var (
	mu       sync.Mutex
	hooks    = map[string]model.Hook{}             // id хука -> хук
	requests = map[string][]model.WebhookRequest{} // id хука -> его запросы
)

type CreateHookResponse struct {
	Status bool   `json:"status"`
	ID     string `json:"id"`
	URL    string `json:"url"`
}

// CreateHook — POST /api/v1/hooks
func CreateHook(c echo.Context) error {
	active := true
	hook := model.Hook{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		IsActive:  &active,
	}

	mu.Lock()
	hooks[hook.ID] = hook
	requests[hook.ID] = []model.WebhookRequest{}
	mu.Unlock()

	url := c.Scheme() + "://" + c.Request().Host + "/wh/" + hook.ID

	return c.JSON(http.StatusCreated, CreateHookResponse{
		Status: true,
		ID:     hook.ID,
		URL:    url,
	})
}

// собирает запрос, сохраняет, отвечает 200.
func CatchHook(c echo.Context) error {
	hookID := c.Param("hookId")
	r := c.Request()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status": false,
			"error":  "cannot read body",
		})
	}

	headers := make(map[string]string, len(r.Header)+1)
	for name, values := range r.Header {
		headers[name] = strings.Join(values, ", ")
	}
	headers["Host"] = r.Host

	req := model.WebhookRequest{
		Method:     r.Method,
		Path:       r.URL.RequestURI(),
		Headers:    headers,
		Body:       string(body),
		IP:         c.RealIP(),
		ReceivedAt: time.Now(),
	}

	mu.Lock()
	_, ok := hooks[hookID]
	if ok {
		requests[hookID] = append(requests[hookID], req)
	}
	mu.Unlock()

	if !ok {
		return c.JSON(http.StatusNotFound, echo.Map{
			"status": false,
			"error":  "hook not found",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": true})
}

func GetRequests(c echo.Context) error {
	hookID := c.Param("id")

	mu.Lock()
	_, ok := hooks[hookID]
	var result []model.WebhookRequest
	if ok {
		src := requests[hookID]
		result = make([]model.WebhookRequest, len(src))
		copy(result, src)
	}
	mu.Unlock()

	if !ok {
		return c.JSON(http.StatusNotFound, echo.Map{
			"status": false,
			"error":  "hook not found",
		})
	}

	return c.JSONPretty(http.StatusOK, result, " ")
}
