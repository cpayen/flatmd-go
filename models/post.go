package models

import "html/template"

// Post is ...
type Post struct {
	Title       string
	URL         string
	Date        string
	Description string
	Content     template.HTML
}
