package controllers

import (
	"github.com/gin-gonic/gin"
	liveKit "github.com/wangu-96/liveKit"
)

func GetLiveKitJoinToken(c *gin.Context) {
	room := c.Query("room")
	identity := c.Query("identity")

	token, err := liveKit.GetJoinToken("your_api_key", "your_api_secret", room, identity)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
