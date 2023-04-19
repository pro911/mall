package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mall/internal/pkg/errno"
	"mall/internal/pkg/response"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = errno.Success
		token := c.GetHeader("token")
		if token == "" {
			code = errno.ErrParam
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				code = errno.ErrAuthFailed
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = errno.ErrInvalidToken
			}
		}

		if code != errno.Success {
			response.ErrorWithMsg(c, code, "")
			c.Abort()
			return
		}
		c.Next()
	}
}

var jwtSecret = []byte(viper.GetString("app.jwt_secret"))

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(Username, Password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Username: Username,
		Password: Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "pro911",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
