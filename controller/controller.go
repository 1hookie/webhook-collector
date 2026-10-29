package controller

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"webhooks/model"
	"webhooks/storage"
)

type Handler struct {
	store   *storage.Store
	baseURL string
}

// конструктор контроллера.
func NewHandler(store *storage.Store, baseURL string) (*Handler, error) {
	switch {
	case store == nil:
		return nil, errors.New("store is nil")
	case baseURL == "":
		return nil, errors.New("baseURL is empty")
	}
	return &Handler{store: store, baseURL: baseURL}, nil
}

type CreateHookResponse struct {
	Status bool   `json:"status"`
	ID     string `json:"id"`
	URL    string `json:"url"`
}

func hookNotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, echo.Map{
		"status": false,
		"error":  "hook not found",
	})
}

// CreateHook — POST /api/v1/hooks
func (h *Handler) CreateHook(c echo.Context) error {
	active := true
	hook := model.Hook{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		IsActive:  &active,
	}

	h.store.CreateHook(hook)

	return c.JSON(http.StatusCreated, CreateHookResponse{
		Status: true,
		ID:     hook.ID,
		URL:    hook.URL(h.baseURL),
	})
}

// собирает запрос, сохраняет, отвечает 200.
func (h *Handler) CatchHook(c echo.Context) error {
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

	if err := h.store.AddRequst(hookID, req); err != nil {
		if errors.Is(err, storage.ErrHookNotFound) {
			return hookNotFound(c)
		}
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{"status": true})
}

func (h *Handler) GetRequests(c echo.Context) error {
	reqs, err := h.store.Requests(c.Param("id"))
	if err != nil {
		if errors.Is(err, storage.ErrHookNotFound) {
			return hookNotFound(c)
		}
		return err
	}
	return c.JSON(http.StatusOK, reqs)
}
