package middleware

import (
	"context"
	"net/http"

	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/config"
	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
            c.Abort()
            return
        }

        ctx := context.Background()
        exists, err := config.RedisClient.Exists(ctx, "blacklist:"+tokenString).Result()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        if exists == 1 {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is in the blacklist"})
            c.Abort()
            return
        }

        claims, err := utils.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        c.Set("userID", claims.Claims.(jwt.MapClaims)["userId"])
        c.Next()
    }
}