package utils

import (
	"fmt"
	"github.com/bbcloudGroup/gothic"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserID 		uint	`json:"user_id"`
	Name 		string 	`json:"name"`
	Mail 		string	`json:"mail"`
	Avatar		string  `json:"avatar"`
	Mobile		string 	`json:"mobile"`
	jwt.StandardClaims
}


func JwtToken(claims *Claims) (signedToken string, success bool) {
	claims.ExpiresAt = time.Now().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(gothic.Config("JwtSecret").(string)))
	if err != nil {
		return
	}
	success = true
	return
}


func JwtValidate(signedToken string) (claims *Claims, success bool) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return []byte(gothic.Config("JwtSecret").(string)), nil
		})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		success = true
		return
	}

	return
}
