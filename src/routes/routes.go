package routes

import (
	"github.com/charoleizer/my-collection/src/domain"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()

	r.GET("/hello_sekai", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": domain.HelloSekai(),
		})

	})

	r.Run()
}
