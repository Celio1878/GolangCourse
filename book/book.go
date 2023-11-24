package book

type IBook interface {
	ShowBooks() ([]Book, error)
	DeleteBook() ([]Book, error)
}

type Book struct {
	Id      int
	Title   string
	Author  string
	Subject string
}

var books = []Book{
	{Id: 1, Title: "The Alchemist", Author: "Paulo Coelho", Subject: "Adventure"},
	{Id: 2, Title: "The Prophet", Author: "Kahlil Gibran", Subject: "Prophet"},
	{Id: 3, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Subject: "Fiction"},
	{Id: 4, Title: "The Odyssey", Author: "Homer", Subject: "Epic"},
	{Id: 5, Title: "The Catcher in the Rye", Author: "J.D. Salinger", Subject: "Fiction"},
}

func (b Book) ShowBooks() ([]Book, error) {
	return books, nil
}

func (b Book) DeleteBook() ([]Book, error) {
	for i, book := range books {
		if book.Id == b.Id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	return books, nil
}
