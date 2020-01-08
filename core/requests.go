package core

import (
	"fmt"
	"gomd/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleHomeRequest(c *gin.Context) {
	books := data.GetBooks()

	c.HTML(http.StatusOK, "home", gin.H{
		"PageTitle":       "Welcome",
		"PageDescription": "Welcome to gomd",
		"Books":           books,
	})
}

func handleBookRequest(c *gin.Context) {
	bookSlug := c.Param("bookSlug")

	book, err := data.GetBook(bookSlug)

	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusNotFound, "error", nil)
		c.Abort()
		return
	}

	posts := data.GetPostsForBook(bookSlug)

	c.HTML(http.StatusOK, "book", gin.H{
		"PageTitle":       book.Title,
		"PageDescription": book.Description,
		"Book":            book,
		"Posts":           posts,
	})
}

func handlePostRequest(c *gin.Context) {
	bookSlug := c.Param("bookSlug")
	postSlug := c.Param("postSlug")

	book, err := data.GetBook(bookSlug)
	post, err := data.GetPost(bookSlug, postSlug)
	bookPosts := data.GetPostsForBook(bookSlug)

	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusNotFound, "error", nil)
		c.Abort()
		return
	}

	c.HTML(http.StatusOK, "post", gin.H{
		"PageTitle":       post.Title,
		"PageDescription": post.Description,
		"Post":            post,
		"Book":            book,
		"BookPosts":       bookPosts,
	})
}
