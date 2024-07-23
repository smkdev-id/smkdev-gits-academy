package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

var db *gorm.DB
var err error

func main() {
	// Initialize Echo
	e := echo.New()

	// Initialize Gorm with SQLite
	db, err = gorm.Open(sqlite.Open("users_management.db"), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal("failed to connect to database")
	}

	// Migrate the User schema
	db.AutoMigrate(&User{})

	e.POST("/register", createUser)
	e.GET("/users", getAllUsers)	
	e.GET("/user/:id/detail", getUserByID)
	e.PUT("/user/:id/update", updateUser)
	e.DELETE("/user/:id/delete", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handlers
func createUser(c echo.Context) error {
	var request User
	if err := c.Bind(&request); err != nil {
		return err
	}

	newUser := User{
		ID:       uuid.New(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully created user",
		"data":    newUser,
	})
}

func getAllUsers(c echo.Context) error {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully retrieved users",
		"data":    users,
	})
}

func getUserByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user ID"})
	}

	var user User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "user not found"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully retrieved user",
		"data":    user,
	})
}

func updateUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user ID"})
	}

	var request User
	if err := c.Bind(&request); err != nil {
		return err
	}

	var user User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "user not found"})
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = request.Password

	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully updated user",
		"data":    user,
	})
}

func deleteUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid user ID"})
	}

	var user User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "user not found"})
	}

	if err := db.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to delete user"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully deleted user",
		"data":    user,
	})
}
