package routes

import "github.com/gin-gonic/gin"

// GetEvents - Retrieves a list of all events
func GetEvents(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "Retrieving Events"})
}
