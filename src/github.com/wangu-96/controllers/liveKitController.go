package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	liveKit "github.com/wangu-96/liveKit"
)

func GetLiveKitJoinToken(c *gin.Context) {
	room := c.Query("room")
	identity := c.Query("identity")

	apiKey := os.Getenv("LIVEKIT_API_KEY")
	apiSecret := os.Getenv("LIVEKIT_API_SECRET")

	token, err := liveKit.GetJoinToken(apiKey, apiSecret, room, identity)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
