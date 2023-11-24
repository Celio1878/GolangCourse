package person

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{name, age}
}

// The better way to work with Receiver Methods is
// Choice either VALUE or POINTER
// Both can cause inconsistency and invalid state

/** In Pointers
* The same object is changed
* Because your address in memory is appointed
 */

/** In Values
* A new object is created
* And allocated in a new position in memory
 */

func (p *Person) ShowPerson() *Person {
	fmt.Println(p)

	return p
}

func (p *Person) ChangePerson(name string) *Person {
	(*p).Name = name

	// Return a value this memory address is pointing at
	fmt.Println("Person -> ", *p)

	// Return a pointer to the memory address
	fmt.Println("Person with pointer -> ", &p)

	return p
}
