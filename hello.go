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

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefic
	case french:
		prefix = frenchHelloPrefic
	}

	return prefix + name
}

func main() {
	fmt.Println(hello("World", "English"))
}
