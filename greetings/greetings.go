package greetings

const englishHelloPrefix = "Hello, "

func Hello(text string) string {
	if text == "" {
		text = "World"
	}

	return englishHelloPrefix + text
}
