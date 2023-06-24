package controller

import (
	"encoding/json"
	"log"
	"mangosteen/config/queries"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/test-go/testify/assert"
)

func TestCreateSession(t *testing.T) {
	teardownTest := SetupTestCase(t)
	defer teardownTest(t)
	w := httptest.NewRecorder()
	sc := SessionController{}
	sc.RegisterRoutes(r.Group("/api"))

	email := "12@qq.com"
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
	log.Println(x)
	bytes, _ := json.Marshal(x)
	req, _ := http.NewRequest("POST", "/api/v1/session", strings.NewReader(string(bytes)))

	r.ServeHTTP(w, req)
	log.Println(w.Body)
	var responseBody struct {
		JWT    string `json:"jwt"`
		UserId int32  `json:"userId"`
	}
	log.Println(&responseBody)
	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Error("jwt is not a string")
	}
	log.Println(w.Body.String())

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, user.ID, responseBody.UserId)

}
