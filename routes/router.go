package routes

import "github.com/gin-gonic/gin"

// SetupRouter - Sets up router to use ginGonic
func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/events", GetEvents)
		v1.GET("/users", GetUsers)
		v1.GET("/colors", GetColors)
	}

	return router
}
