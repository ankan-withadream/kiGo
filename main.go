package main

import (
	"net/http"

	"github.com/ankan-withadream/kiGo/config"
	"github.com/gin-gonic/gin"
)

// Demo struct of Book
type Book struct {
	ID     string `json:"id" binding:"required"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// A Demo Store for string all the books
var books = map[string]Book{
	"1": {ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	"2": {ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	"3": {ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

// main is the entry point of the application.
// It initializes the Gin router and sets up the HTTP endpoints for CRUD operations on books.
func main() {
	// Initialize of Gin router
	router := gin.Default()

	// GET /book: Retrieves all books from the map.
	// It responds with a JSON array of books.
	router.GET("/book", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, books)
	})

	/*
			 POST /book: Adds a new book to the map.
		     It expects a JSON payload containing the book's ID, title, and author.
		     If the payload is invalid, it responds with a JSON object containing an error message.
		     Otherwise, it adds the new book to the map and responds with a JSON object containing the new book.
	*/

	router.POST("/book", func(c *gin.Context) {

		var newBook Book

		// Call BindJSON to bind the received JSON to newBook structure.
		if err := c.BindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Add the new book to the map.
		books[newBook.ID] = newBook
		c.IndentedJSON(http.StatusCreated, newBook)

	})

	/*
	   PUT /book: Updates an existing book in the map.
	   It expects a JSON payload containing the book's ID, title, and author.
	   If the payload is invalid, it responds with a JSON object containing an error message.
	   Otherwise, it updates the corresponding book in the map and responds with a JSON object containing the updated book
	*/
	router.PUT("/book", func(c *gin.Context) {
		var updatedBook Book

		if err := c.BindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		books[updatedBook.ID] = updatedBook

		c.IndentedJSON(http.StatusAccepted, updatedBook)

	})

	/*
	   PUT /book: Updates an existing book in the map.
	   It expects a JSON payload containing the book's ID, title, and author.
	   If the payload is invalid, it responds with a JSON object containing an error message.
	   Otherwise, it updates the corresponding book in the map and responds with a JSON object containing the updated book
	*/
	router.PATCH("/book", func(c *gin.Context) {
		var updatedBook Book

		if err := c.BindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		books[updatedBook.ID] = updatedBook

		c.IndentedJSON(http.StatusAccepted, updatedBook)

	})

	/*

	   DELETE /book: Deletes a book from the map.
	   It expects a query parameter containing the book's ID.
	   If the ID is not provided, it responds with a JSON object containing an error message.
	   Otherwise, it deletes the corresponding book from the map and responds with a JSON object containing a success message.
	*/
	router.DELETE("/book", func(c *gin.Context) {
		id := c.Query("id")
		delete(books, id)
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "book deleted",
		})
	})

	// Gin Server run on mentioned Port
	router.Run(config.PORT)
}
