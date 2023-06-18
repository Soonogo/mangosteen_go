package controller

import (
	"context"
	"mangosteen/config"
	"mangosteen/internal/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	// viper.Set("email.smtp.host", "localhost")
	// viper.Set("email.smtp.port", 1025)
	r = gin.Default()
	config.LoadAppConfig()
	database.Connect()
	vc := ValidationCodeController{}
	vc.RegisterRoutes(r.Group("/api"))

	email := "tttsongen@gmail.com"
	c := context.Background()
	q := database.NewQuery()
	count1, _ := q.CountValidationCodes(c, email)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/validation_codes", strings.NewReader(`{"email":"`+email+`"}`))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	count2, _ := q.CountValidationCodes(c, email)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, count1+1, count2)
}
