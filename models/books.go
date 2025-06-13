package models

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	Genre    string `json:"genre"`
}

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
