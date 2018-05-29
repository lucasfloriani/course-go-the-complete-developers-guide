package main

import "fmt"

// Definition to what other types need
// to be a type bot
type bot interface {
	getGreeting() string
}

// They are of type bot too because they have
// function getGreeting() string
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}
