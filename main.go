package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const AssetBasePath = "./public/assets"
const PartParam = "part"
const NameParam = "name"

const TemplatesBasePath = "./templates/*"

func main() {
	router := gin.Default()

	router.LoadHTMLGlob(TemplatesBasePath) // Loading templates folder
	router.Static("/public", "./public")

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{})

	})

	router.GET("/time", func(c *gin.Context) {
		currentTime := time.Now()
		c.String(http.StatusOK, currentTime.String())
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
