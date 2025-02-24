package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", callback)
	router.Run(":8080")
}

func callback(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping-pong and ding-dong",
	})
}
