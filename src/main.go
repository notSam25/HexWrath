// src/main.go

package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.SetFuncMap(template.FuncMap{})
	router.LoadHTMLGlob("src/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.tmpl", gin.H{
			"title":   "Welcome",
			"message": "Hello, World!",
		})
	})
	router.Run(":8080")
}
