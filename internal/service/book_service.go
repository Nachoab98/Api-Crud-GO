package service

import (
	"errors"
	"goApi/internal/model"
	"goApi/internal/store"
)
/*type Logger interface {
	Log(msg string, error string)
}*/
type Service struct {
	store store.Store
	//logger Logger
}
func New(s store.Store) *Service {
	return &Service{
		store: s,
		//logger: nil,
	}
}
func (s *Service) GetAllBooks() ([]model.Book, error) {
	//s.logger.Log("We are fetching all books", "")
	books, err:= s.store.GetAllBooks()
	if err != nil {
		//s.logger.Log("Error fetching books %v\n ", err.Error())
		return nil, err
	}
	return books, nil
}
func (s *Service) GetBookById(id int) (model.Book, error) {
	return s.store.GetBookByID(id)
}
func (s *Service) CreateBook(book model.Book) (model.Book, error) {
	if book.Title == "" {
		return model.Book{}, errors.New("title is required")
		}
	if book.Author == "" {
		return model.Book{}, errors.New("author is required")
	}
	return s.store.CreateBook(book)
}
func (s *Service) UpdateBook(id int, book model.Book) (model.Book, error) {
	if book.Title == "" {
		return model.Book{}, errors.New("title is required")
	}
	return s.store.UpdateBook(id, book)
}
func (s *Service) DeleteBook(id int)  error {
	return s.store.DeleteBook(id)
}