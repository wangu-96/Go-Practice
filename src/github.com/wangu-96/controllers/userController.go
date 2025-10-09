package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wangu-96/JWT"
	"github.com/wangu-96/initializers"
	"github.com/wangu-96/models"
)

// func UsersCreate(c *gin.Context) {

// 	var Body struct {
// 		Name     string `json:"name"`
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	if err := c.BindJSON(&Body); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user := models.User{Name: Body.Name, Email: Body.Email, Password: Body.Password}
// 	result := initializers.DB.Create(&user)

// 	if result.Error != nil {
// 		log.Fatal(result.Error)
// 	}

// 	c.JSON(200, gin.H{
// 		"message": "User created successfully",
// 		"user":    user,
// 	})

// }

func UsersCreate(c *gin.Context) {
	var Body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&Body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:     Body.Name,
		Email:    Body.Email,
		Password: Body.Password,
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// ✅ Generate JWT token
	token, err := JWT.CreateToken(user.Email) // or user.Name / user.ID
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with user and token
	c.JSON(200, gin.H{
		"message": "User created successfully",
		"user":    user,
		"token":   token,
	})
}
