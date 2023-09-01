package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go_echo_api/utils"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if err := HelloHandler(c); err != nil {
		t.Error("HelloHandler() error = ", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	
	expectedJSON := utils.GenerateExpectedJSON("Hello World")
	assert.Equal(t, string(expectedJSON), rec.Body.String())
}
