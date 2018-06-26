package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/qrcode", func(c *gin.Context) {
		content := c.DefaultQuery("content", "https://github.com/NeroCube")
		if pic, err := qrcode.Encode(content, qrcode.Medium, 256); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.Data(http.StatusOK, "image/png", pic)
		}
	})
	router.Run()
}
