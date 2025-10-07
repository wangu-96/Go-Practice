package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wangu-96/initializers"
	"github.com/wangu-96/models"
)

func PostsCreate(c *gin.Context) {

	// 1. Create a variable to hold the input
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	// 2. Bind the request JSON into `body`
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 3. Use that data to create your model
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	c.JSON(200, gin.H{
		"message": "Post created successfully",
		"post":    post,
	})

}

func PostIndex(c *gin.Context) {

	var posts = []models.Post{}
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	c.JSON(200, gin.H{
		"message": "Posts fetched successfully",
		"posts":   posts,
	})

}

func ShowPost(c *gin.Context) {

	id := c.Param("id")
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	}

	c.JSON(200, gin.H{
		"message": "Post fetched successfully",
		"post":    post,
	})

}

func UpdatePost(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
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

	// Refresh the post to get updated values
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"message": "Post updated successfully",
		"post":    post,
	})

}

func DeletePost(c *gin.Context) {

	id := c.Param("id")

	var post models.Post
	// if err := initializers.DB.First(&post, id).Error; err != nil {
	// 	c.JSON(404, gin.H{"error": "Post not found"})
	// 	return
	// }

	if err := initializers.DB.Unscoped().Delete(&post, id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}
