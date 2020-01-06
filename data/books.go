package data

import (
	"encoding/json"
	"gomd/models"
	"io/ioutil"
	"log"
)

// GetBooks is...
func GetBooks() []models.Book {
	var books []models.Book

	files, err := ioutil.ReadDir("./content/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			bookSlug := file.Name()
			datafile, err := ioutil.ReadFile("./content/" + bookSlug + "/data.json")

			//todo: log error?
			if err == nil {
				book := models.Book{URL: bookSlug}
				json.Unmarshal(datafile, &book)
				books = append(books, book)
			}
		}
	}

	return books
}

// GetBook is...
func GetBook(bookSlug string) (models.Book, error) {

	datafile, err := ioutil.ReadFile("./content/" + bookSlug + "/data.json")
	if err != nil {
		return models.Book{}, err
	}

	book := models.Book{URL: bookSlug}
	json.Unmarshal(datafile, &book)

	return book, nil
}
