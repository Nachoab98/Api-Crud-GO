package transport

import (
	"goApi/internal/service"
	"net/http"
)

type BookHandler struct {
	service *service.Service
}

func New(s *service.Service) *BookHandler{
	return &BookHandler{
		service: s,
	}
} 

func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		books, err:= h.service.GetAllBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)

	case http.MethodPost:
		var book model.Book
		if err:= json.NewDecoder(r.Body).Decode(&book); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		created, err:= h.service.CreateBook(book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(created)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}		
}