package core

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Init app
func Init() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{", "}}")

	r.Use(static.Serve("/assets", static.LocalFile("./theme/assets", false)))

	r.LoadHTMLGlob("./theme/templates/*.tpl.html")

	// Routes
	r.GET("/", func(c *gin.Context) { handleHomeRequest(c) })
	r.GET("/:bookSlug", func(c *gin.Context) { handleBookRequest(c) })
	r.GET("/:bookSlug/:postSlug", func(c *gin.Context) { handlePostRequest(c) })

	r.Run()
}
