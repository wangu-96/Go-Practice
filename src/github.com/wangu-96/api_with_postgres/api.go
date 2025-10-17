package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wangu-96/controllers"
	"github.com/wangu-96/initializers"
	"github.com/wangu-96/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnnectToDatabase()
}

func main() {
	r := gin.Default()

	//User routes
	r.POST("/users", controllers.UsersCreate)
	r.GET("/login", controllers.UsersLogin)

	//This is just to test getting a token for conference call api
	r.GET("/livekit/token", controllers.GetLiveKitJoinToken)

	protected := r.Group("/")
	protected.Use(middleware.RequireAuth())
	{

		//Post routes
		protected.POST("/post", controllers.PostsCreate)
		protected.GET("/posts", controllers.PostIndex)
		protected.GET("/posts/:id", controllers.ShowPost)
		protected.PUT("/posts/:id", controllers.UpdatePost)
		protected.DELETE("/posts/:id", controllers.DeletePost)

	}
	log.Println("Server running at http://localhost:3000")

	r.Run(":3000") // listen and serve on
}
