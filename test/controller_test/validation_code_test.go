package controller_test

import (
	"context"
	"fmt"
	"log"
	"mangosteen/internal/database"
	"mangosteen/internal/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	viper.Set("email.smtp.host", "localhost")
	viper.Set("email.smtp.port", 1025)
	r := router.New()
	email := "tttsongen@gmail.com"
	c := context.Background()
	q := database.NewQuery()
	count1, _ := q.CountValidationCodes(c, email)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/validation_code", strings.NewReader(`{"email":"`+email+`"}`))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	count2, _ := q.CountValidationCodes(c, email)

	fmt.Println(req)

	log.Println(count1, count2, "----------")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, count1+1, count2)
}
