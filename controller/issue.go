package controller

import (
	"ginx/cds"
	"ginx/config"
	"ginx/utility"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Issue issue jwt token
func Issue(c *gin.Context) {

	claims := cds.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "ginx",
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
		UID: uuid.NewV4().String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(config.App.Secret))
	if err != nil {
		utility.Logger().Error("generate signed string error, " + err.Error())
		c.String(http.StatusInternalServerError, "generate sined string error")
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": signed,
	})
	return
}

// CheckIssue authorize issue
func CheckIssue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "authorized",
		"uid": c.GetString("uid"),
	})
	return
}
