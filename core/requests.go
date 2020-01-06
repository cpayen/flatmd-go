package core

import (
	"fmt"
	"gomd/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleHomeRequest(c *gin.Context) {
	books := data.GetBooks()

	c.HTML(http.StatusOK, "index.tpl.html", gin.H{
		"Books": books,
	})
}

func handleBookRequest(c *gin.Context) {
	bookSlug := c.Param("bookSlug")

	book, err := data.GetBook(bookSlug)

	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusNotFound, "error.tpl.html", nil)
		c.Abort()
		return
	}

	posts := data.GetPostsForBook(bookSlug)

	c.HTML(http.StatusOK, "book.tpl.html", gin.H{
		"Book":  book,
		"Posts": posts,
	})
}

func handlePostRequest(c *gin.Context) {
	bookSlug := c.Param("bookSlug")
	postSlug := c.Param("postSlug")

	post, err := data.GetPost(bookSlug, postSlug)

	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusNotFound, "error.tpl.html", nil)
		c.Abort()
		return
	}

	c.HTML(http.StatusOK, "post.tpl.html", gin.H{
		"Post": post,
	})
}
