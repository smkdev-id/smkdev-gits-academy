package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/config"
	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/utils"
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

        fmt.Println(claims)


        userIDFloat, ok := claims["userID"].(float64)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            c.Abort()
            return
        }

        userID := strconv.FormatUint(uint64(userIDFloat), 10)
        c.Set("userID", userID)
        fmt.Println(userID)
        c.Next()
    }
}

func AdminOnly() gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.GetString("userID")
        if userID == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
            c.Abort()
            return
        }

        userIDUint, err := strconv.ParseUint(userID, 10, 32)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
            c.Abort()
            return
        }

        isAdmin, err := utils.IsUserAdmin(uint(userIDUint))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        if !isAdmin {
            c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
            c.Abort()
            return
        }

        c.Next()
    }
}