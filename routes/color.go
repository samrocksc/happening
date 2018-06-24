package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/samrocksc/gofunzi/db"
	m "github.com/samrocksc/gofunzi/models"
)

// GetColors - gets all users
func GetColors(c *gin.Context) {
	db := db.GetDB()

	var results []m.Color
	if err := db.Find(&results).Error; err != nil {
		fmt.Println(err)
	}

	c.JSON(200, results)
}
