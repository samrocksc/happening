package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/samrocksc/gofunzi/db"
	m "github.com/samrocksc/gofunzi/models"
)

// GetUsers - gets all users
func GetUsers(c *gin.Context) {
	db := db.GetDB()

	var users []m.User
	if err := db.Find(&users).Error; err != nil {
		fmt.Println(err)
	}

	c.JSON(200, users)
}
