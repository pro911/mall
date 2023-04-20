package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mall/internal/pkg/errno"
	"mall/internal/pkg/response"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			response.ErrorWithMsg(c, errno.ErrMustLogin, "")
			c.Abort()
			return
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				response.ErrorWithMsg(c, errno.ErrAuthFailed, err.Error())
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				response.ErrorWithMsg(c, errno.ErrInvalidToken, "")
				c.Abort()
				return
			}
			c.Set("user_id", claims.UserID)
			zap.L().Info("解析token", zap.Int("user_id", claims.UserID))
		}

		c.Next()
	}
}

var jwtSecret = []byte(viper.GetString("app.jwt_secret"))

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(UserID int, Username, Password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		UserID:   UserID,
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
