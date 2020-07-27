package main

import (
	"fmt"
	"time"
)
var wg = sync.WaitGroup{}  
func portal1(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "portal 1"
}
func portal2(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "portal 2"
}
func fibonocci(c chan int, num int) {
	mul := 1
	for i := 1; i <= num; i++ {
		mul *= i
	}
	c <- mul
}
func main() {
	wg.Add(3)
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan int)
	go portal1((c1))
	go portal2(c2)
	go fibonocci(c3, 5)
	for {
		count := 0
		select {
		case v1 := <-c1:
			fmt.Println(v1)
			count++
		case v2 := <-c2:
			count++
			fmt.Println(v2)
		case v3 := <-c3:
			fmt.Println(v3)
			count++
		default:
			if count == 3 {
				break
			}
		}

	}
}
