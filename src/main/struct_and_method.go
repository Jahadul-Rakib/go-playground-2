package main

import "fmt"

type Book struct {
	bookName   string
	bookPrice  float64
	bookWriter string
}



func main() {

	var bookOne Book = Book{"bangla", 123.40, "rakib"}
	var bookTwo Book = Book{"math", 123.40, "rakib"}
	var bookThree Book = Book{"english", 123.40, "rakib"}

	fmt.Print(bookOne, bookTwo, bookThree)
	//book := Book{"physics", 110, "Jahadul"}


	book := bookTwo.changeObjectValue("physics", 110, "Jahadul")
	fmt.Print(book)

}


func (b Book) changeObjectValue(bookName string, bookPrice float64, bookWriter string) Book {
	b.bookName = bookName
	b.bookPrice = bookPrice
	b.bookWriter = bookWriter

	return b
}