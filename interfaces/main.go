package main

import "fmt"

// englishBot represents a bot object that speaks english
type englishBot struct{}

// spanishBot represents a bot object that speaks spanish
type spanishBot struct{}

// bot represents an interface for bots
type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// getGreeting returns the greeting for the englishBot
func (eb englishBot) getGreeting() string {
	return "Hello!"
}

// getGreeting returns the greeting for the spanishBot
func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

// printGreeting prints out the greeting of the bot
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
