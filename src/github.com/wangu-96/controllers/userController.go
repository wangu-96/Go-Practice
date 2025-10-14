package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wangu-96/JWT"
	"github.com/wangu-96/initializers"
	"github.com/wangu-96/models"
	"golang.org/x/crypto/bcrypt"
)

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

	// ✅ Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Name:     Body.Name,
		Email:    Body.Email,
		Password: string(hashedPassword),
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// ✅ Generate JWT token
	token, err := JWT.CreateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with user and token
	c.JSON(200, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
		"token": token,
	})
}

func UsersLogin(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON input
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Find the user by email
	var user models.User
	if err := initializers.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare the provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(500, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := JWT.CreateToken(user.ID, user.Email)
	if err != nil {
		log.Println("Token generation error:", err)
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with token and user info
	c.JSON(200, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
		"token": token,
	})
}
