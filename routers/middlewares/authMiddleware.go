package middlewares

import (
	"time"

	"hito/configs"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	UserId       string
	PasswordHash string
}

var jwtKey string
var jwtLifetime time.Duration

func init() {
	jwtKey = configs.GeneralConf.GetString("app.jwt.key")
	jwtLifetime = configs.GeneralConf.GetDuration("app.jwt.lifetime")
}

// func AuthJWT(c *gin.Context) {
// 	auth := c.GetHeader("Authorization")
// 	splits := strings.Split(auth, "Bearer ")
// 	if len(splits) < 2 {
// 		resp.Err(c, errs.InvalidToken)
// 		c.Abort()
// 		return
// 	}

// 	token := splits[1]
// 	user := validToken(token)
// 	if user == nil {
// 		resp.Err(c, errs.InvalidToken)
// 		c.Abort()
// 		return
// 	}

// 	accountInfo, err := controllers.GetAccountInfo(user)
// 	if err == nil {
// 		account.SetAccountInfo(c, accountInfo)
// 	}

// 	setUser(c, user)
// 	setSessionToken(c, token)

// 	c.Next()
// }

// func CreateToken(user *models.HmxUser) (string, error) {
// 	claims := Claims{
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(jwtLifetime).Unix(),
// 		},
// 		UserId:       user.UserId,
// 		PasswordHash: user.PasswordHash,
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	return token.SignedString([]byte(jwtKey))
// }

// func validToken(tokenString string) *models.HmxUser {
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
