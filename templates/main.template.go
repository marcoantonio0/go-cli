package templates

var MainTemplate string = `
package main

import (
	"github.com/gin-gonic/gin"
	"{{.ProjectName}}/app/config"
	route "{{.ProjectName}}/app/routes"
)

func main() {
	r := gin.Default()

	config.InitDatabase()

	route.UserRoutes(r.Group("/user"))

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"name":    "{{.ProjectName}}",
			"version": 1.0,
		})
	})

	r.Run(":3000")
}
`
