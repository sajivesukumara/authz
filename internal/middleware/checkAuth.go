package middleware

import (
	"fmt"
	"authz/internal/initializers"
	"authz/internal/models"
	"net/http"
	"os"
	"strings"
	"time"
    "strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
        fmt.Printf("Error  => %v\n", err)
		return "", err
	}

    claims, ok := token.Claims.(jwt.MapClaims)
    fmt.Printf("Claims => %v", claims)

	if ok && token.Valid {
		username := claims["id"].(float64)
		return strconv.FormatFloat(username, 'f', -1, 64), nil
	}

	return "", fmt.Errorf("invalid token")
}

func CheckAuth(c *gin.Context){
    authHeader := c.GetHeader("Authorization")

    if authHeader == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    authTokens := strings.Split(authHeader, " ")
    if len(authTokens) != 2 || authTokens[0] != "Bearer" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is invalid"})
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    tokenString := authTokens[1]

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }   
        return []byte(os.Getenv("SECRET")), nil
    })

    if err != nil || !token.Valid {
        fmt.Printf("Invalid Token ERROR: [%s]\n", err)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    if float64(time.Now().Unix()) > claims["exp"].(float64) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    var user models.User
    initializers.DB.Where("id=?", claims["id"]).First(&user)
    if user.ID == 0 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    c.Set("currentUser", user)
    c.Next()
}