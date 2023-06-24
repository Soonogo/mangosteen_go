package controller

import (
	"crypto/rand"
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/email"

	"github.com/gin-gonic/gin"
)

type ValidationCodeController struct{}

func (ctrl *ValidationCodeController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	v1.POST("/validation_codes", ctrl.Create)
}

func (ctrl *ValidationCodeController) Create(c *gin.Context) {
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(400, "参数错误")
		return
	}
	q := database.NewQuery()

	str := generateDigists(6)

	log.Println("[CreateValidationCode]", body)
	vc, err := q.CreateValidationCode(c, queries.CreateValidationCodeParams{
		Email: body.Email,
		Code:  str,
	})
	if err != nil {
		// TODO 没有做校验
		c.Status(400)
		return
	}
	if err := email.SendValidationCode(vc.Email, vc.Code); err != nil {
		log.Println("[SendValidationCode fail]", err)
		c.String(500, "发送失败")
		return
	}
	c.Status(200)
}

func (ctrl *ValidationCodeController) Destroy(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ValidationCodeController) Update(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ValidationCodeController) Get(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ValidationCodeController) GetPaged(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

// helper function
func generateDigists(len int) string {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		log.Println("[CreateValidationCode]", err)
	}
	digits := make([]byte, len)
	for i := range b {
		log.Println(b[i])
		digits[i] = b[i]%10 + 48
	}
	return string(digits)
}
