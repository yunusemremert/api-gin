package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Author   string `json:"author,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

var books = []book{
	{ID: "1", Title: "Uçurtma Avcısı", Author: "Khaled Hosseini", Quantity: 1},
	{ID: "2", Title: "Şeker Portakalı", Author: "Jose Mauro De Vasconcelos", Quantity: 4},
	{ID: "3", Title: "Simyacı", Author: "Paulo Coelho", Quantity: 2},
}

func main() {
	router := gin.Default()

	router.GET("/api/books", getBooks)
	router.GET("/api/book/:id", getBook)
	router.POST("/api/book", createBook)
	router.PATCH("/api/book", updateBook)

	router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {
	id := c.Param("id")

	book, err := bookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})

		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func bookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)
}

func updateBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})

		return
	}

	book, err := bookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})

		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})

		return
	}

	book.Quantity -= 1

	c.IndentedJSON(http.StatusOK, book)
}
