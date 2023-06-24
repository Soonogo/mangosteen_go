package controller

import (
	"context"
	"mangosteen/config"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/middleware"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
	q *queries.Queries
	c context.Context
)

func SetupTestCase(t *testing.T) func(t *testing.T) {
	r = gin.Default()
	config.LoadAppConfig()
	database.Connect()
	r.Use(middleware.Me())
	q = database.NewQuery()
	c = context.Background()

	if err := q.DeleteAllUsers(c); err != nil {
		t.Fatal(err)
	}
	return func(t *testing.T) {
		database.Close()
	}

}
