package main

import "fmt"

const (
	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	frenchHelloPrefix  = "Ca va "

	spanish = "spanish"
	french  = "French"
)

func Hello(name, language string) string {
	if len(name) == 0 {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("world", ""))
}
