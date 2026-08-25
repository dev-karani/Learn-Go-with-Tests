package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	swahiliHelloPrefix = "Wasalaimkum, "

	swahili = "Swahili"
	spanish = "Spanish"
	french  = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case swahili:
		prefix = swahiliHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Elodie", "Swahili"))
}
