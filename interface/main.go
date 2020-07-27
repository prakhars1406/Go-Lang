package main

import (
	"fmt"
	"os"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	arg := os.Args
	fmt.Println(arg[1])
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
	return "Hola"
}

//func printGreeting(eb englishBot) {
//	fmt.Println(eb.getGreeting)
//}

//func printGreeting(sb spanishBot) {
//	fmt.Println(sb.getGreeting)
//}
