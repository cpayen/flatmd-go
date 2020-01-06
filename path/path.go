package path

import (
	"fmt"
)

var (
	content string = "./content/"
	book    string = "./content/%s/"
	post    string = "./content/%s/%s/"
)

// GetContentPath is..
func GetContentPath() string {
	return content
}

// GetBookPathFor is..
func GetBookPathFor(bookSlug string) string {
	return fmt.Sprintf(book, bookSlug)
}

// GetPostPathFor is..
func GetPostPathFor(bookSlug string, postSlug string) string {
	return fmt.Sprintf(post, bookSlug, postSlug)
}
