package data

import "dev1/models"

var Users []models.User

var Books = []models.Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marchel", Quantity: 2, Genre: "Action"},
	{ID: "2", Title: "The Great Gatsby", Author: "Scott", Quantity: 5, Genre: "Action"},
	{ID: "3", Title: "War and Peace", Author: "Leo Mass", Quantity: 6, Genre: "RomCom"},
	{ID: "4", Title: "In Search ", Author: "Marchel", Quantity: 2, Genre: "RomCom"},
}

var Authors = []models.Author{
	{ID: "1", Name: "Marchel"},
	{ID: "2", Name: "Scott"},
	{ID: "3", Name: "Leo Mass"},
}
