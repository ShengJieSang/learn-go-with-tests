package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const franch = "Franch"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {

	switch language {

	case franch:
		prefix = frenchHelloPrefix

	case spanish:
		prefix = spanishHelloPrefix

	default:
		prefix = englishHelloPrefix
	}

	return
}
