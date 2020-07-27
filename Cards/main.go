package main

import (
	"fmt"
	"os"
)

func main() {
	//var card string = "ace of spaids"
	//card := "ace of spades"
	//card = "Five of diamonds"
	//	card := newCard()
	//cards := []string{"xyz", "zyx"}

	//cards := deck{"xyz", "zyx"}
	//cards = append(cards, "abc")
	//fmt.Println(cards)

	//	for i, card := range cards {
	//	fmt.Println(i, card)
	//}

	cards := newDeck()
	cards.print()

	card1, card2 := deal(cards, 5)
	fmt.Println(card1)
	fmt.Println(card2)

	greeting := "Hi There"
	fmt.Println([]byte(greeting))

	fmt.Println(cards.toString())

	cards.saveToFile("my_cards")

	newDeckFromFile("my_cards").print()

	cards.shuffle()
	cards.print()
	createFile("Myt_test_file")

}
func newCard() string {
	return "Five of diamonds"
}

func createFile(fileName string) {
	fs, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	_, err = fs.Write([]byte("Hello world"))
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	defer fs.Close()
}
