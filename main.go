package main

import (
	"dev1/api"
	"dev1/handler"
	"dev1/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/signup", handler.SignUp)                                   // Add a new User
	router.POST("/login", handler.Login)                                     // user login
	router.POST("/books", middleware.AuthMiddleware, api.CreateBook)         // add a new book
	router.GET("/books/:id", middleware.AuthMiddleware, api.BookbyId)        // find book by id
	router.PATCH("/checkout", middleware.AuthMiddleware, api.CheckOutBook)   // take a book from library
	router.PATCH("/checkin", middleware.AuthMiddleware, api.CheckInBook)     // submit a book in library
	router.PATCH("/updatebook", middleware.AuthMiddleware, api.UpdateBook)   //update a book with json
	router.PATCH("/deleteBook", middleware.AuthMiddleware, api.DeleteBook)   //delete a book by id
	router.GET("/author/:name", middleware.AuthMiddleware, api.BookbyAuthor) // find book by  author name
	router.GET("/genre/:name", middleware.AuthMiddleware, api.BookbyGenre)   // find book by Genre
	router.GET("/authorname/:id", middleware.AuthMiddleware, api.FindAuthor) // id to author name
	router.Run("localhost:8080")
}
