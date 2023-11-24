package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type Result struct {
	Message string
	Err     error
}

func Hello(name string) Result {
	if name == "" {
		errorMessage := "empty name"
		return Result{"", errors.New(errorMessage)}
	}

	greets := greetings{
		"Hello %v",
		"Greeting to see you, %v",
		"Hail, %v! Well met !!!",
	}

	// Return the range of elements from 1 index to 3
	fmt.Println(greets[1:3])

	randomMessage := greets.randomGreeting()
	message := fmt.Sprintf(randomMessage, name)

	// Saving in local file
	err := greets.savingLocalFile("greetings.txt")
	if err != nil {
		return Result{"", err}
	}

	// Reading from file
	greetsFromFile, err := greetingsFromFile("greetings.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		// Get out from program e show the error code
		os.Exit(400)

		return Result{"", err}
	}
	fmt.Println("Reading data from file ->", greetsFromFile)

	return Result{message, nil}
}

// greetings Receiver Function type
// When created a custom type, it is possible to create a function receiver
// It works like a method
type greetings []string

func (g greetings) randomGreeting() string {
	greets := append(g, "HI %v")

	return greets[rand.Intn(len(greets))]
}

func (g greetings) savingLocalFile(filename string) error {
	fmt.Println("Saving in file...")

	return os.WriteFile(filename, []byte(g.randomGreeting()), 0644)

}

func greetingsFromFile(filename string) (greetings, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return greetings{string(data)}, nil
}
