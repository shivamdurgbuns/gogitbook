package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefic = "Hola, "
const frenchHelloPrefic = "Bonjour, "

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefic + name
	}

	if language == french {
		return frenchHelloPrefic + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(hello("World", "English"))
}
