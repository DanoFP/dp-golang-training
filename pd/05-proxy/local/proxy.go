package local

import "github.com/DanoFP/dp-golang-training/pd/05-proxy/book"

type Proxy interface {
	GetByID(ID uint) book.Book
	GetAll() book.Books
}
