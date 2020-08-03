package main

import (
	"fmt"
	"time"
	"sync"
	
)
var goRoutine sync.WaitGroup
func one(ch chan string){
	defer goRoutine.Done()
	for i:=0;i<10;i++{
		time.Sleep(1*time.Second)
	}	
	ch <-"done one"
}
func two(ch chan string){
 	defer goRoutine.Done()
	for i:=0;i<12;i++{
		time.Sleep(1*time.Second)
	}	
	ch <-"done two"

}
func main() {
	ch:=make(chan string,20)
	ch2:=make(chan string,20)
	goRoutine.Add(1)
	go one(ch)
	goRoutine.Add(1)
	go two(ch2)
	fmt.Println(<-ch)
	fmt.Println(<-ch2)
	goRoutine.Wait()
}
