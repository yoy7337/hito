package v1

import (
	"hito/controllers"
	"hito/models"
	"hito/routers/errs"
	"hito/routers/resp"

	"github.com/gin-gonic/gin"
)

func MountUserApis(publicR *gin.RouterGroup, privateR *gin.RouterGroup) {
	publicR.POST("/login", login)
	publicR.POST("/signup", signUp)
}

// @Summary 登入
// @Tags User 使用者
// @version 1.0
// @Accept       json
// @Produce      json
// @Success 200 object interface{}
// @Router /user/login [post]
func login(c *gin.Context) {
	var opts controllers.LoginOpts
	if err := c.ShouldBindJSON(&opts); err != nil {
		c.Error(errs.InternalError.AppendMsg(err.Error()))
		return
	}

	var user models.User
	if err := opts.Login(&user); err != nil {
		c.Error(errs.Msg(errs.CanNotLogin, err.Error()))
		return
	}

	resp.JSON(c, gin.H{
		"message": "login",
		"status":  "ok",
	})
}

// @Summary 註冊
// @Tags User 使用者
// @version 1.0
// @Accept       json
// @Produce      json
// @Success 200 object interface{}
// @Router /user/signup [post]
func signUp(c *gin.Context) {
	var opts controllers.SignUpOpts
	if err := c.ShouldBindJSON(&opts); err != nil {
		c.Error(errs.InternalError.AppendMsg(err.Error()))
		return
	}

	var user models.User
	if err := opts.SignUp(&user); err != nil {
		c.Error(errs.Msg(errs.CanNotCreateUser, err.Error()))
		return
	}

	resp.JSON(c, user)
}
