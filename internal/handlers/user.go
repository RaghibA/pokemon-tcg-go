package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/RaghibA/pokemon-tcg-go/internal/database"
	"github.com/RaghibA/pokemon-tcg-go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Username string
	Password string
	Email    string
}

//* Helper Functions
func validateUsername(u string) bool { //! move to handler func
	user := models.User{}
	database.PokeDb.Db.First(&user, "username = ?", u)

	return user.ID == 0
}

func validateEmail(email string) bool {
	user := models.User{}
	database.PokeDb.Db.First(&user, "email = ?", email)

	return user.ID == 0
}

//* Create User Handler
func CreateUserHandler(c *gin.Context) {

	var newUser RegisterBody

	if c.Bind(&newUser) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"code":    4000001,
		})

		return
	}

	if !validateUsername(newUser.Username) { // validate the username
		c.JSON(http.StatusConflict, gin.H{ // this should be returned in the helper function
			"message": "An account with this username already exists",
			"code":    4090001,
		})

		return
	}

	if !validateEmail(newUser.Email) {
		c.JSON(http.StatusConflict, gin.H{
			"message": "An account with this email address already exists",
			"code":    4090010,
		})

		return
	}

	// Add user to database
	db := true
	if db != true {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Unable to process request",
			"code":    4090005,
		})
	}

	// If everything is okay, return 200 & account info
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Unable to process request",
			"code":    4090005,
		})

		return
	}

	user := models.User{
		Username: newUser.Username,
		Password: string(hash),
		Email:    newUser.Email,
	}

	database.PokeDb.Db.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message":  "Account created",
		"username": newUser.Username,
		"email":    newUser.Email,
	})
}

//* Login User Handler
func LoginUserHandler(c *gin.Context) {
	// Get credentials from body
	var newUser RegisterBody

	if c.Bind(&newUser) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"code":    4000001,
		})

		return
	}

	// Look up user
	user := models.User{}
	database.PokeDb.Db.First(&user, "username = ?", newUser.Username)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"code":    4040001,
		})

		return
	}

	// Compare pass with pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newUser.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid username or password",
			"code":    4010001,
		})

		return
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		log.Println("Failed to generate JWT")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal service error",
			"code":    "5000001",
		})

		return
	}

	// Return token
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

//* Delete User Handler
func DeleteUserHandler(c *gin.Context) {
	user, _ := c.Get("user")

	//! Need to delete/overwrite cookie here too

	database.PokeDb.Db.Delete(&user, user)
	c.JSON((http.StatusOK), gin.H{})
}

//* Authenticate User Handler
func AuthenticateUserHandler(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Authenticated",
		"data":    user,
	})
}

// Logout user handler
func LogoutUserHandler(c *gin.Context) {
	// send new cookie to delete/overwrite old one
}
