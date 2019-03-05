package main

import "github.com/gin-gonic/gin"

func main() {
	server := Routes()
	server.Run(":8080")
}

func Routes() *gin.Engine {
	application := gin.Default()
	application.GET("/health-check", func(c *gin.Context) {
		c.String(200,"Service OK! =]")
	})

	return application
}

