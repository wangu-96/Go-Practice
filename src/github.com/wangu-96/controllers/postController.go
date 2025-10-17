package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wangu-96/initializers"
	"github.com/wangu-96/models"
)

// creates a logged in users post (tested and working)
func PostsCreate(c *gin.Context) {
	// 1. Get the authenticated user's ID from context
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDInterface.(uint)

	// 2. Create a variable to hold the input
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	// 3. Bind the request JSON into `body`
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 4. Use that data to create your model and associate it with the user
	post := models.Post{
		Title:  body.Title,
		Body:   body.Body,
		UserID: userID, // associate post with logged-in user
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post created successfully",
		"post":    post,
	})
}

func PostIndex(c *gin.Context) {
	// Get the currently logged-in user ID (assuming it's stored in context by middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var posts []models.Post
	result := initializers.DB.Where("user_id = ?", userID).Find(&posts)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(500, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Posts fetched successfully",
		"posts":   posts,
	})
}

func ShowPost(c *gin.Context) {
	// Get the post ID from URL params
	id := c.Param("id")

	// Get the currently logged-in user's ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var post models.Post
	// Find the post by ID **and** user_id
	result := initializers.DB.Where("id = ? AND user_id = ?", id, userID).First(&post)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post not found or you don't have access"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post fetched successfully",
		"post":    post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	// Get logged-in user's ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Bind input JSON
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Find the post with ID AND user_id
	var post models.Post
	if err := initializers.DB.Where("id = ? AND user_id = ?", id, userID).First(&post).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found or you don't have access"})
		return
	}

	// Update fields
	updated := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	if err := initializers.DB.Model(&post).Updates(updated).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update post"})
		return
	}

	// Refresh the post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"message": "Post updated successfully",
		"post":    post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	// Get logged-in user's ID from context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	// Find the post with ID AND user_id
	var post models.Post
	if err := initializers.DB.Where("id = ? AND user_id = ?", id, userID).First(&post).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found or you don't have access"})
		return
	}

	// Delete the post
	if err := initializers.DB.Unscoped().Delete(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}
