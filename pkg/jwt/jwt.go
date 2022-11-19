package jwt

import (
	"bluebell/pkg/response"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var salt = []byte("salt")       // 加密盐
var LoginExpireTime = time.Hour //超时时间

// 自定义加密结构体
type loginClaims struct {
	Username string `json:"username"`
	UserId   string `json:"userid"`
	jwt.RegisteredClaims
}

// type RegisteredClaims struct {
// 	// the `iss` (Issuer) claim. ""  签发人
// 	Issuer string `json:"iss,omitempty"`

// 	// the `sub` (Subject) claim.  主题
// 	Subject string `json:"sub,omitempty"`

// 	// the `aud` (Audience) claim. 受众
// 	Audience ClaimStrings `json:"aud,omitempty"`

// 	// the `exp` (Expiration Time) claim.  过期时间
// 	ExpiresAt *NumericDate `json:"exp,omitempty"`

// 	// the `nbf` (Not Before) claim. 生效时间
// 	NotBefore *NumericDate `json:"nbf,omitempty"`

// 	// the `iat` (Issued At) claim. 签发时间
// 	IssuedAt *NumericDate `json:"iat,omitempty"`

// 	// the `jti` (JWT ID) claim. jwtID号
// 	ID string `json:"jti,omitempty"`
// }

func GenerateJwt(username, userid string) (string, error) {
	claims := loginClaims{
		username,
		userid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(LoginExpireTime)),
			Issuer:    "tomato-admin", // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(salt)
}
func Parsejwt(tokenstring string) (*loginClaims, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &loginClaims{}, func(t *jwt.Token) (interface{}, error) {
		return salt, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*loginClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 中间件
func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authString := c.Request.Header.Get("token")
		if authString == "" {
			response.ResponseJSON(c, response.CodeTokenEmpty, nil)
			c.Abort()
			return
		}
		mc, err := Parsejwt(authString)
		if err != nil {
			response.ResponseJSON(c, response.CodeTokenInvalid, map[string]any{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("username", mc.Username)
		c.Set("userid", mc.UserId)
		c.Next()
	}
}
