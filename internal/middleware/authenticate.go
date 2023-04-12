package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/RaghibA/pokemon-tcg-go/internal/database"
	"github.com/RaghibA/pokemon-tcg-go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(c *gin.Context) {
	// Store bearer token from request
	tokenString := c.Request.Header["X-Auth-Token"][0]
	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode & validate token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Failed to authenticate token: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Validate claims
		exp := claims["exp"].(float64)
		if float64(time.Now().Unix()) > exp {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find user
		user := models.User{}
		database.PokeDb.Db.First(&user, claims["subject"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}

//! No longer using cookies for auth

// func AuthenticateCookie(c *gin.Context) {
// 	// Get cookie off request
// 	tokenString, err := c.Cookie("Authorization")
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 	}

// 	// decode & validate cookie
// 	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Failed to authenticate token: %v", t.Header["alg"])
// 		}

// 		return []byte(os.Getenv("SECRET")), nil
// 	})

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

// 		// Validate claims
// 		exp := claims["exp"].(float64)
// 		if float64(time.Now().Unix()) > exp {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}

// 		// Find user
// 		user := models.User{}
// 		database.PokeDb.Db.First(&user, claims["subject"])

// 		if user.ID == 0 {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}

// 		c.Set("user", user)

// 	} else {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 	}

// 	c.Next()
// }
