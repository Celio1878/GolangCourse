package main

import (
	"GolangCourse/book"
	"GolangCourse/checker"
	"GolangCourse/greeting"
	"GolangCourse/person"
	"fmt"
	"log"
	"rsc.io/quote"
)

func main() {
	// Show logs only if occurs any errors
	log.SetPrefix("Greetings Logs\n")
	/** Flag Values
	0: nothing;
	1: date;
	2: time;
	3: date and time;
	8: file;
	19: date, time and package name;
	*/
	log.SetFlags(19)

	// About Greeting Module
	println("Greeting Module")
	const name string = "Master"
	message := greeting.Hello(name)

	if message.Err != nil {
		log.Fatal(message.Err)
		return
	}

	fmt.Println(message.Message)

	// Get a byte slice of the message
	fmt.Println([]byte(message.Message))

	fmt.Println(quote.Go())

	// About Person Module
	println("\nPerson Module")
	newPerson := person.NewPerson("Celio", 30)
	fmt.Println("Showing New Person")
	newPerson.ShowPerson()

	changedPerson := newPerson.ChangePerson("Swangle")
	fmt.Println("Updating Person Name", changedPerson)

	changedPersonMemoryAddress := &changedPerson
	println("Updating Person Memory Address", changedPersonMemoryAddress)
	newPerson.ShowPerson()

	// About Map
	println("\nInterface Module")
	var ib book.IBook = book.Book{Id: 1, Title: "Harry Potter", Author: "J.K Rowling", Subject: "Fantasy"}

	bs, err := ib.ShowBooks()
	if err != nil {
		return
	}
	fmt.Println("Books are: ", bs)

	b, e := ib.DeleteBook()
	if e != nil {
		return
	}
	fmt.Println("\nUpdated Books are: ", b)

	// About Channel
	println("\nChannel Module")
	checker.WebsiteChecker()

}

// Value types - int, float, string, bool, struct (Use POINTERS to change it in a function)
// Reference types - slice, map, channel, pointer, function (Use REFERENCE so, don't need to use POINTERS to change it in a function)
// Interface types - interface{}
// Array types - [n]T
// Slice types - []T
// Map types - map[K]V
// Channel types - chan T
// Pointer types - *T
// Function types - func("Afghanistan")
// Struct types - struct { Name string; Age int }
