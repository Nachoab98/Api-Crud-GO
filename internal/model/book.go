package model

type Book struct {
	ID int  `json:"id"`   //para traducir a archivo json/XML
	Title string  `json:"title"`
	Author string  `json:"author"` 
}