package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishKeyWord     = "es"
	frenchKeyWord      = "fr"
)

func Hello(word, language string) string {
	if word == "" {
		return englishHelloPrefix + "world"
	}

	var helloTraduction string

	switch language {
	case spanishKeyWord:
		helloTraduction = "Ola, "
	case frenchKeyWord:
		helloTraduction = "Salut, "
	default:
		helloTraduction = "Hello, "
	}

	return helloTraduction + word
}

func main() {
	fmt.Println(Hello("world", "fr"))
}
