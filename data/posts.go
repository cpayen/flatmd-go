package data

import (
	"encoding/json"
	"gomd/models"
	"gomd/path"
	"html/template"
	"io/ioutil"
	"log"

	"gopkg.in/russross/blackfriday.v2"
)

// GetPostsForBook is...
func GetPostsForBook(bookSlug string) []models.Post {
	var posts []models.Post

	files, err := ioutil.ReadDir(path.GetBookPathFor(bookSlug))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			postSlug := file.Name()
			datafile, err := ioutil.ReadFile(path.GetPostPathFor(bookSlug, postSlug) + "data.json")

			//todo: log error?
			if err == nil {
				post := models.Post{URL: bookSlug + "/" + postSlug}
				json.Unmarshal(datafile, &post)
				posts = append(posts, post)
			}
		}
	}

	return posts
}

// GetPost is...
func GetPost(bookSlug string, postSlug string) (models.Post, error) {

	datafile, err := ioutil.ReadFile(path.GetPostPathFor(bookSlug, postSlug) + "data.json")
	if err != nil {
		return models.Post{}, err
	}

	contentfile, err := ioutil.ReadFile(path.GetPostPathFor(bookSlug, postSlug) + "content.md")
	if err != nil {
		return models.Post{}, err
	}

	content := template.HTML(blackfriday.Run([]byte(contentfile)))

	post := models.Post{URL: bookSlug + "/" + postSlug, Content: content}
	json.Unmarshal(datafile, &post)

	return post, nil
}
