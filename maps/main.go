package main

import "fmt"

func main() {
	//colors := map[string]string{
	//	"red":   "00x",
	//	"green": "01x",
	//}

	//var colors map[string]string

	colors := make(map[int]string)

	colors[10] = "002x"
	colors[20] = "003x"
	colors[30] = "004x"
	//delete(colors, 10)

	for key, value := range colors {
		fmt.Println(key, value)
	}
	//	fmt.Println(colors)
}
