package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(text string) string {
	if text == "" {
		text = "World"
	}

	return englishHelloPrefix + text
}

func main() {
	fmt.Println(Hello("world"))
}
