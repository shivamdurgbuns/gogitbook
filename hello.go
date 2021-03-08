package main

import (
	"fmt"
)

//Setting the constants for the program

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

/*
Hello function will return Hello World to the programmer by default when run.
It will also return Hello, <name> ; if name parameter is passed to the function.
It also has functionality to return Hello. <name> in Spanish and French. If the language is also passed
along with the name.
*/
func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

/*
greetingPrefix function will set the Prefix value of the greeting according to the language passed to program.
*/
func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", "English"))
}
