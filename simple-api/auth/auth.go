package auth

import (
	"go1/simple-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

const (
	USER     = "admin"
	PASSWORD = "Pasword123"
	SECRET   = "secret"
)

func Login(c *gin.Context) {
	var user models.Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request!",
		})
	}

	if user.Username != USER {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid username!",
		})
		c.Abort()
	} else if user.Password != PASSWORD {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password!",
		})
		c.Abort()
	}

	// token
	claim := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		Issuer:    "test",
		IssuedAt:  time.Now().Unix(),
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := sign.SignedString([]byte(SECRET))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	if user.Username == USER && user.Password == PASSWORD {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login success!",
			"token":   token,
		})
	}
}
