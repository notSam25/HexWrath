// src/main.go

package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Set up a template engine using html/template
	router.SetFuncMap(template.FuncMap{
		// Add any custom template functions here if needed
	})
	router.LoadHTMLGlob("src/templates/*") // Load templates from the "templates" directory

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.tmpl", gin.H{
			"title":   "Welcome",
			"message": "Hello, World!",
		})
	})
	router.Run(":8080")
}
