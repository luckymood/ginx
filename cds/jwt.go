package cds

import (
	"github.com/dgrijalva/jwt-go"
)

// JwtClaims jwt claims
type JwtClaims struct {
	jwt.StandardClaims
	UID string `json:"uid"`
}
