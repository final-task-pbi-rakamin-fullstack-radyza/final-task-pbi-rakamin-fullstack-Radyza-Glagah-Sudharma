package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"project-final/app"
	"project-final/helpers"
)

type UserController struct {
	DB *gorm.DB
}

func (ctrl UserController) RegisterUser(c *gin.Context) {
	var user app.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	hashedPassword, _ := helpers.HashPassword(user.Password)
	user.Password = hashedPassword

	if result := ctrl.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

func (ctrl UserController) LoginUser(c *gin.Context) {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
        return
    }

    var user app.User
    result := ctrl.DB.Where("email = ?", credentials.Email).First(&user)
    if result.Error != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if !helpers.CheckPasswordHash(credentials.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, err := helpers.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func (ctrl UserController) UpdateUser(c *gin.Context) {
    var userUpdates app.User
    userID := c.Param("id")

    if err := c.ShouldBindJSON(&userUpdates); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
        return
    }

    result := ctrl.DB.Model(&app.User{}).Where("id = ?", userID).Updates(userUpdates)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (ctrl UserController) DeleteUser(c *gin.Context) {
    userID := c.Param("id") 

    result := ctrl.DB.Delete(&app.User{}, userID)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No user found with provided ID"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
