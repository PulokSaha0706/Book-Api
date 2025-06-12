package main

import (
	"errors"
	//"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	Genre    string `json:"genre"`
}

type author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marchel", Quantity: 2, Genre: "Action"},
	{ID: "2", Title: "The Great Gatsby", Author: "Scott", Quantity: 5, Genre: "Action"},
	{ID: "3", Title: "War and Peace", Author: "Leo Mass", Quantity: 6, Genre: "RomCom"},
	{ID: "4", Title: "In Search ", Author: "Marchel", Quantity: 2, Genre: "RomCom"},
}

var authors = []author{
	{ID: "1", Name: "Marchel"},
	{ID: "2", Name: "Scott"},
	{ID: "3", Name: "Leo Mass"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// helper to find a book by it's id
func getBookbyID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found")
}

func bookbyId(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookbyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	//fmt.Println(newBook.Title)

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func updateBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	book, err := getBookbyID(newBook.ID)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book.Title = newBook.Title
	book.Author = newBook.Author
	book.Quantity = newBook.Quantity

	c.IndentedJSON(http.StatusOK, books)

}

func checkOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missong id in query parameter"})
		return
	}

	book, err := getBookbyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book is not currently available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}

func checkInBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id in query parameter"})
		return
	}

	book, err := getBookbyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}

func deleteBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id in query parameter"})
		return
	}

	book, err := getBookbyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	for i, b := range books {
		if b.ID == book.ID {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, books)
			return
		}
	}

}

func bookbyAuthor(c *gin.Context) {
	name := c.Param("name")

	var newBookList []book

	for _, b := range books {
		if b.Author == name {
			newBookList = append(newBookList, b)
		}
	}

	c.IndentedJSON(http.StatusOK, newBookList)
}

func bookbyGenre(c *gin.Context) {
	name := c.Param("name")

	var newBookList []book

	for _, b := range books {
		if b.Genre == name {
			newBookList = append(newBookList, b)
		}
	}

	c.IndentedJSON(http.StatusOK, newBookList)
}

func findAuthor(c *gin.Context) {
	id := c.Param("id")

	for _, b := range authors {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Authort not found"})

}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)            // see total book list
	router.POST("/books", createBook)         // add a new book
	router.GET("/books/:id", bookbyId)        // find book by id
	router.PATCH("/checkout", checkOutBook)   // take a book from library
	router.PATCH("/checkin", checkInBook)     // submit a book in library
	router.PATCH("/updatebook", updateBook)   //update a book with json
	router.PATCH("/deleteBook", deleteBook)   //delete a book by id
	router.GET("/author/:name", bookbyAuthor) // find book by  author name
	router.GET("/genre/:name", bookbyGenre)   // find book by Genre
	router.GET("/authorname/:id", findAuthor) // id to author name
	router.Run("localhost:8080")
}
