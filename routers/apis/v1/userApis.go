package v1

import (
	"hito/routers/resp"

	"github.com/gin-gonic/gin"
)

func MountUserApis(publicR *gin.RouterGroup, privateR *gin.RouterGroup) {
	publicR.POST("/login", login)
	publicR.POST("/signup", signUp)
}

// @Summary 註冊
// @Tags User 使用者
// @version 1.0
// @Accept       json
// @Produce      json
// @Success 200 object interface{}
// @Router /user/signup [post]
func signUp(c *gin.Context) {
	resp.JSON(c, gin.H{
		"message": "signUp",
		"status":  "ok",
	})
}

// @Summary 登入
// @Tags User 使用者
// @version 1.0
// @Accept       json
// @Produce      json
// @Success 200 object interface{}
// @Router /user/login [post]
func login(c *gin.Context) {
	resp.JSON(c, gin.H{
		"message": "login",
		"status":  "ok",
	})
}
