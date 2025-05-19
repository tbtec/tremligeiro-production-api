package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

// Mock controller implementing IController
type mockController struct {
	lastRequest httpserver.Request
}

func (m *mockController) Handle(ctx context.Context, req httpserver.Request) httpserver.Response {
	m.lastRequest = req
	return httpserver.Response{
		Code: 200,
		Body: map[string]string{"msg": "ok"},
	}
}

func TestAdapt(t *testing.T) {
	app := fiber.New()
	mockCtrl := &mockController{}
	app.Post("/test/:id", adapt(mockCtrl))

	payload := map[string]string{"foo": "bar"}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/test/123?x=1", bytes.NewReader(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Check if controller received correct request
	assert.Equal(t, "/test/123", mockCtrl.lastRequest.Path)
	assert.Equal(t, "POST", mockCtrl.lastRequest.Method)
	assert.Equal(t, "123", mockCtrl.lastRequest.Params["id"])
	assert.Equal(t, "1", mockCtrl.lastRequest.Query["x"])
	assert.JSONEq(t, string(body), string(mockCtrl.lastRequest.Body))
}
