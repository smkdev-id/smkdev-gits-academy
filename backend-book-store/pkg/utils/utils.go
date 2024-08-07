package utils

import (
	"errors"
	"time"

	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/config"
	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("secret")

func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return jwtSecret, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    } else {
        return nil, errors.New("invalid token")
    }
}

func IsUserAdmin(userID uint) (bool, error) {
    var user models.User
    if err := config.DB.Where("id = ?", userID).Preload("Roles").First(&user).Error; err != nil {
        return false, err
    }
    for _, role := range user.Roles {
        if role.Name == "admin" {
            return true, nil
        }
    }
    return false, nil
}