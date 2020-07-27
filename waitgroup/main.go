package main

import (
	"fmt"
	"sync"
	"time"
)

func Calcualte(c chan string, w *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	c <- "Calculate"
	w.Done()
}
func Develope(c chan string, w *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	c <- "Develope"
	w.Done()
}

func main() {
	var wg sync.WaitGroup
	c1 := make(chan string)
	c2 := make(chan string)
	wg.Add(2)
	go Calcualte(c1, &wg)
	go Develope(c2, &wg)
	fmt.Println(<-c1)
	fmt.Printf(<-c2)
	wg.Wait()

}
