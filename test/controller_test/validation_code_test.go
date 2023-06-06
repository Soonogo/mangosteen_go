package controller_test

import (
	"fmt"
	"mangosteen/internal/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	r := router.New()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/validation_code", strings.NewReader(`{"email":"hello"}`))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	fmt.Println(req)

	assert.Equal(t, 200, w.Code)
}
