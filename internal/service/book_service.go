package service
import (
	"goApi/internal/model"
	"goApi/internal/store"
)
type Service struct {
	store store.Store
}
func New(s store.Store) *Service {
	return &Service{
		store: s,
	}
}
func (s *Service) GetAllBooks() ([]model.Book, error) {
	return s.store.GetAllBooks()
}
func (s *Service) GetBookById(id int) (model.Book, error) {
	return s.store.GetBookByID(id)
}
func (s *Service) CreateBook(book model.Book) (model.Book, error) {
	return s.store.CreateBook(book)
}

