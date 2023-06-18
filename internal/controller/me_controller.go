package controller

import (
	"mangosteen/internal/database"
	"mangosteen/internal/jwt_helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type MeController struct{}

func (ctrl *MeController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	v1.GET("me", ctrl.Get)
}

// GetMe godoc
//
//	@Summary	获取当前用户
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	api.GetMeResponse
//	@Failure	401	{string}	JWT为空	|	无效的JWT
//	@Router		/api/v1/me [get]
func (ctrl *MeController) Get(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if len(auth) < 8 {
		c.String(401, "unauthorized")
		return
	}
	jwtString := auth[7:]
	t, err := jwt_helper.Parse(jwtString)
	if err != nil {
		c.String(401, "unauthorized")
		return
	}
	m, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		c.String(401, "unauthorized")
		return
	}
	userID, ok := m["user_id"].(float64)
	if !ok {
		c.String(401, "unauthorized")
		return

	}
	userIDInt := int32(userID)
	if err != nil {
		c.String(401, "无效的JWT")
		return
	}
	q := database.NewQuery()
	userIDInt32 := int32(userIDInt)
	user, err := q.FindUser(c, userIDInt32)
	if err != nil {
		c.String(401, "无效的JWT")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"resource": user,
	})

}

func (ctrl *MeController) Create(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *MeController) Destroy(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *MeController) Update(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *MeController) GetPaged(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
