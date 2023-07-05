package v1

import (
	"hito/controllers"
	"hito/models"
	"hito/routers/errs"
	"hito/routers/middleware"
	"hito/routers/resp"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TokenResp struct {
	Token string `json:"token"` // login token...
}

func MountUserApis(publicR *gin.RouterGroup, privateR *gin.RouterGroup) {
	publicR.POST("/login", login)
	publicR.POST("/signup", signUp)
	privateR.GET("/tap", tap)
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

	// login and get user
	var user models.User
	if err := opts.Login(&user); err != nil {
		c.Error(errs.Msg(errs.CanNotLogin, err.Error()))
		return
	}

	// create token (JWT)
	token, err := middleware.CreateJWTToken(&user)
	if err != nil {
		c.Error(errs.InternalError.AppendMsg(err.Error()))
		return
	}

	log.Debugf("token: %s", token)

	resp.JSON(c, TokenResp{Token: token})
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

// @Summary 確認是否登入
// @Tags User 使用者
// @version 1.0
// @Accept       json
// @Produce      json
// @Success 200 object interface{}
// @Router /user/tap [get]
func tap(c *gin.Context) {
	resp.JSON(c, gin.H{"msg": "ok"})
}
