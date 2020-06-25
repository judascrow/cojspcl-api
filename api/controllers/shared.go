package controllers

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/judascrow/gomiddlewares/jwt"
	"golang.org/x/crypto/bcrypt"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ClaimsOwner(c *gin.Context, slug string) bool {

	claims := jwt.ExtractClaims(c)

	if slug == claims["slug"].(string) || ClaimsIsAdmin(claims) {
		return true
	}
	return false
}

func ClaimsIsAdmin(claims jwt.MapClaims) bool {

	if int(claims["roleId"].(float64)) == 1 {
		return true
	}

	return false
}
