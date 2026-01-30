package store
import (
	"database/sql"
	"goApi/internal/model"
)
type Store interface  {
	GetAllBooks()([]model.Book, error)
	GetBookByID(id int)(model.Book, error)
	CreateBook(book model.Book)(model.Book, error)
	UpdateBook(id int, book model.Book)(model.Book, error)
	DeleteBook(id int) error
}
type store struct{
	db *sql.DB
}
func NewStore(db *sql.DB) Store {
	return &store{
		db: db,
	}

}
func (s *store) GetAllBooks()([]model.Book, error){
	q := `SELECT * FROM books`         //definimos la query
	rows, err := s.db.Query(q)			//ejecutamos la query sobre la "store"
	if err != nil {				
		return nil, err
	}
	defer rows.Close()    //para cerrar la conexion despues de que termine de ejecutar

	var books []model.Book				//slice que permite seguir agregando libros a diferencia del array donde hay que determinar un largo fijo
	for rows.Next(){
		var b model.Book  	 
		err := rows.Scan(&b.ID, &b.Title, &b.Author)			
		if err != nil {
			return nil, err
	}
	books = append(books, b)
}
return books, nil
}
func (s *store) GetBookByID(id int)(model.Book, error){
	q:= `SELECT * FROM books WHERE id = ?`
	var b model.Book
	err := s.db.QueryRow(q, id).Scan(&b.ID, &b.Title, &b.Author)

	if err != nil {
		return model.Book{}, err  // struct vacío, no nil
	}
	return b, nil
}
func (s *store) CreateBook(book model.Book)(model.Book, error){
	q:= `INSERT INTO books (title, author) VALUES (?, ?)`
	res, err := s.db.Exec(q, book.Title, book.Author)
	if err != nil {
		return model.Book{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return model.Book{}, err
	}
	book.ID = int(id)			
	return book, nil
}
func (s *store) UpdateBook(id int, book model.Book)(model.Book, error){
	q:= `UPDATE books SET title = ?, author = ? WHERE id = ?`
	_, err := s.db.Exec(q, book.Title, book.Author, id)
	if err != nil {
		return model.Book{}, err
	}
	book.ID = id
	return book, nil
}
func (s *store) DeleteBook(id int) error{
q:= `DELETE FROM books WHERE id = ?`
_, err := s.db.Exec(q, id)
if err != nil {
	return  err
}
return  nil
}