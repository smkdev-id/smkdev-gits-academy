package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("YNTKTS")

func JWTAuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		tokenString := c.GetHeader("Authorization")

		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Hapus prefix "Bearer " dari token string
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		log.Printf("Token: %+v", tokenString)

		// Parse token dan validasi
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			log.Printf("Token: %+v", token)
			log.Printf("Token key: %+v", jwtKey)

			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			log.Println("Error parsing token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Ambil klaim dari token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Printf("Token claims: %+v", claims)
			log.Printf("Token claims: %+v", claims["role"].(string))

			userRole, roleOk := claims["role"].(string)
			userID, idOk := claims["user_id"].(string)

			// Cek apakah klaim role dan user_id valid
			if !roleOk || !idOk {
				log.Printf("Role mismatch: required %s, but got %s", requiredRole, userRole)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				c.Abort()
				return
			}

			// Simpan user_id dan role ke context untuk digunakan di handler lain
			c.Set("user_id", userID)
			c.Set("role", userRole)

			// Cek apakah peran pengguna sesuai dengan yang dibutuhkan
			if userRole != requiredRole {
				log.Println("Invalid token claims or token not valid")
				c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Lanjut ke handler berikutnya
		c.Next()
	}
}
