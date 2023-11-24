package checker

import (
	"fmt"
	"net/http"
	"time"
)

func WebsiteChecker() chan string {
	links := []string{
		"https://www.google.com",
		"https://www.quii.gitbook.io/learn-go-with-tests.com",
		"https://www.golang.com",
		"https://www.amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		fmt.Println("Value inside the channel: ", <-c)
		go func(link string) {
			fmt.Println("Sleeping 5 seconds")
			const fiveSeconds = 5 * time.Second
			time.Sleep(fiveSeconds)
			checkLink(link, c)
		}(l)
	}

	return c
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down: ", err)
		c <- "down"
		return
	}

	fmt.Println(link, "is up!")
	c <- "up"
}
