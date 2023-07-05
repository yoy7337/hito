package middleware

import (
	"fmt"
	"strings"
	"time"

	"hito/configs"
	"hito/models"
	"hito/routers/errs"
	"hito/routers/resp"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	jwt.StandardClaims
	UserOId      string
	PasswordSalt string
}

var jwtKey string
var jwtLifetime time.Duration

func init() {
	jwtKey = configs.GeneralConf.GetString("app.jwt.key")
	jwtLifetime = configs.GeneralConf.GetDuration("app.jwt.lifetime")
}

func AuthJWT(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	splits := strings.Split(auth, "Bearer ")
	if len(splits) < 2 {
		resp.Err(c, errs.InvalidToken)
		c.Abort()
		return
	}

	token := splits[1]
	user := validJWTToken(token)
	if user == nil {
		resp.Err(c, errs.InvalidToken)
		c.Abort()
		return
	}

	log.Debugf("user auth: %v", user)

	// accountInfo, err := controllers.GetAccountInfo(user)
	// if err == nil {
	// 	account.SetAccountInfo(c, accountInfo)
	// }

	// setUser(c, user)
	// setSessionToken(c, token)

	c.Next()
}

func CreateJWTToken(user *models.User) (string, error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtLifetime).Unix(),
			Id:        user.Id,
		},
		UserOId:      user.ID.Hex(),
		PasswordSalt: user.PasswordSalt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtKey))
}

func validJWTToken(tokenString string) *models.User {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// check signing method
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtKey), nil
	})

	// can not validate token...
	if err != nil || !token.Valid {
		log.Warn("invalid token")
		return nil
	}

	oId, err := primitive.ObjectIDFromHex(claims.UserOId)
	if err != nil {
		return nil
	}

	var user models.User
	opts := models.FindOpts{
		Decode: &user,
	}

	if err := opts.FindOneByOId(user.ModelName(), &oId); err != nil {
		log.Warn("user not found")
		return nil
	}

	// check password salt, passwordSalt will be changed when user change password
	if user.PasswordSalt != claims.PasswordSalt {
		log.Warn("password salt changed")
		return nil
	}

	return &user
}

// TODO: login with long-session
// 	session, err := models.GetSession(tokenString)

// 	// can not parse token...
// 	if err != nil {
// 		return nil
// 	}

// 	user, err := models.GetUser(session.UserOId)
// 	if err != nil {
// 		return nil
// 	}

// 	return user
// }

// func setUser(c *gin.Context, user *models.HmxUser) {
// 	if user != nil {
// 		c.Set("user", user)
// 	}
// }

// func GetUser(c *gin.Context) *models.HmxUser {
// 	user, exists := c.Get("user")

// 	if !exists {
// 		return nil
// 	}

// 	return user.(*models.HmxUser)
// }

// func setSessionToken(c *gin.Context, token string) {
// 	c.Set("sessionToken", token)
// }

// func GetSessionToken(c *gin.Context) string {
// 	return c.GetString("sessionToken")
// }
