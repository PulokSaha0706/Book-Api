package api

import (
	"dev1/data"
	"dev1/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Books)
}

// helper to find a book by it's id
func GetBookbyID(id string) (*models.Book, error) {
	for i, b := range data.Books {
		if b.ID == id {
			return &data.Books[i], nil
		}
	}

	return nil, errors.New("Book not found")
}

func BookbyId(c *gin.Context) {
	id := c.Param("id")
	book, err := GetBookbyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	//fmt.Println(newBook.Title)

	data.Books = append(data.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func UpdateBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	book, err := GetBookbyID(newBook.ID)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book.Title = newBook.Title
	book.Author = newBook.Author
	book.Quantity = newBook.Quantity

	c.IndentedJSON(http.StatusOK, data.Books)

}

func CheckOutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missong id in query parameter"})
		return
	}

	book, err := GetBookbyID(id)

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

func CheckInBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id in query parameter"})
		return
	}

	book, err := GetBookbyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}

func DeleteBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id in query parameter"})
		return
	}

	book, err := GetBookbyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	for i, b := range data.Books {
		if b.ID == book.ID {
			data.Books = append(data.Books[:i], data.Books[i+1:]...)
			c.IndentedJSON(http.StatusOK, data.Books)
			return
		}
	}

}

func BookbyAuthor(c *gin.Context) {
	name := c.Param("name")

	var newBookList []models.Book

	for _, b := range data.Books {
		if b.Author == name {
			newBookList = append(newBookList, b)
		}
	}

	c.IndentedJSON(http.StatusOK, newBookList)
}

func BookbyGenre(c *gin.Context) {
	name := c.Param("name")

	var newBookList []models.Book

	for _, b := range data.Books {
		if b.Genre == name {
			newBookList = append(newBookList, b)
		}
	}

	c.IndentedJSON(http.StatusOK, newBookList)
}

func FindAuthor(c *gin.Context) {
	id := c.Param("id")

	for _, b := range data.Authors {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Authort not found"})

}
