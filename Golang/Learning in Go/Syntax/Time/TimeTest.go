package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"1": "123",
			"2": "333",
		})
	})
	r.Any("/any", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"1": "123",
			"2": "333",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
