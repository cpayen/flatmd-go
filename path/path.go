package path

import (
	"fmt"
)

var (
	content string = "./content/pages/"
	book    string = "./content/pages/%s/"
	post    string = "./content/pages/%s/%s/"
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
