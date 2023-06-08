package controller_test

import (
	"context"
	"encoding/json"
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/test-go/testify/assert"
)

func TestCreateSession(t *testing.T) {
	r := router.New()
	w := httptest.NewRecorder()
	email := "1@qq.com"
	code := "1234"
	q := database.NewQuery()
	c := context.Background()
	if _, err := q.CreateValidationCode(c, queries.CreateValidationCodeParams{
		Email: email, Code: code,
	}); err != nil {
		log.Fatalln(err)
	}
	x := gin.H{
		"email": email,
		"code":  code,
	}
	bytes, _ := json.Marshal(x)
	req, _ := http.NewRequest("POST", "/api/v1/session", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	log.Println(w.Body)
	var responseBody struct {
		JWT string `json:"jwt"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Error("jwt is not a string")
	}
	log.Println(w.Body.String())

	assert.Equal(t, 2020, w.Code)
}
