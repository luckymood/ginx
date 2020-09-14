package middleware

import (
	"ginx/cds"
	"ginx/config"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JwtAuthorize jwt 身份验证
func JwtAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.String(http.StatusUnauthorized, "there is no authorization")
			c.Abort()
			return
		}
		token, _ := jwt.ParseWithClaims(tokenString, &cds.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.App.Secret), nil
		})
		if claims, ok := token.Claims.(*cds.JwtClaims); ok && token.Valid {
			c.Set("uid", claims.UID)
		} else {
			c.String(http.StatusUnauthorized, "invalid claims")
			c.Abort()
			return
		}
		c.Next()
	}
}
