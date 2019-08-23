package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Athens!",
		})
	})

	if err := r.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
