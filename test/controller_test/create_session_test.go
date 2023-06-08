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

var (
	r *gin.Engine
	q *queries.Queries
	c context.Context
)

func setupTest(t *testing.T) func(t *testing.T) {
	r = router.New()
	q = database.NewQuery()
	c = context.Background()
	if err := q.DeleteAllUsers(c); err != nil {
		t.Fatal(err)
	}
	return func(t *testing.T) {
		database.Close()
	}

}

func TestCreateSession(t *testing.T) {
	teardownTest := setupTest(t)
	defer teardownTest(t)
	w := httptest.NewRecorder()

	email := "1@qq.com"
	code := "1234"
	if _, err := q.CreateValidationCode(c, queries.CreateValidationCodeParams{
		Email: email, Code: code,
	}); err != nil {
		log.Fatalln(err)
	}
	user, err := q.CreateUser(c, email)
	if err != nil {
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
		JWT    string `json:"jwt"`
		UserId int32  `json:"userId"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Error("jwt is not a string")
	}
	log.Println(w.Body.String())

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, user.ID, responseBody.UserId)

}
