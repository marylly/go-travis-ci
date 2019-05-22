package main

import "github.com/gin-gonic/gin"
import "os"

func main() {
	port := getPort()
	server := Routes()
	server.Run(port)
}

func Routes() *gin.Engine {
	application := gin.Default()
	application.GET("/health-check", func(c *gin.Context) {
		c.String(200,"Service OK! =]")
	})

	return application
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":3000"
}
