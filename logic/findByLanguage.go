package logic

import (
	"../template"
	"strings"
	"../data"
)
func FindByLanguage(keyword string) template.Books {

	var foundBooks template.Books
	foundBooks.Books = make([]template.Book, 0)

	for i:=0;i< len(data.BOOKS.Books);i++ {

		book:= data.BOOKS.Books[i]

		if book != (template.Book{}) && 
			strings.Contains(strings.ToUpper(book.Language),strings.ToUpper(keyword)){

			foundBooks.Books = append(foundBooks.Books,book)
			
		}
	}
	
	return foundBooks
}
