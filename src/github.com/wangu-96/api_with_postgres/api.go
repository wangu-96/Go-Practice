package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wangu-96/controllers"
	"github.com/wangu-96/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnnectToDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/post", controllers.PostsCreate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.ShowPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	log.Println("Server running at http://localhost:3000")

	r.Run(":3000") // listen and serve on

}
