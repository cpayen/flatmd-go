package core

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func createRenderer() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()
	renderer.AddFromFiles("home", "./theme/templates/layout.tpl.html", "./theme/templates/home.tpl.html")
	renderer.AddFromFiles("book", "./theme/templates/layout.tpl.html", "./theme/templates/book.tpl.html")
	renderer.AddFromFiles("post", "./theme/templates/layout.tpl.html", "./theme/templates/post.tpl.html")
	renderer.AddFromFiles("error", "./theme/templates/layout.tpl.html", "./theme/templates/error.tpl.html")
	return renderer
}

// Init app
func Init() {
	router := gin.Default()
	router.HTMLRender = createRenderer()
	router.Use(gin.Logger())
	router.Use(static.Serve("/assets", static.LocalFile("./theme/assets", false)))
	router.Use(static.Serve("/", static.LocalFile("./content/pages", false)))

	// Routes
	router.GET("/", func(c *gin.Context) { handleHomeRequest(c) })
	router.GET("/:bookSlug", func(c *gin.Context) { handleBookRequest(c) })
	router.GET("/:bookSlug/:postSlug", func(c *gin.Context) { handlePostRequest(c) })

	router.Run()
}
