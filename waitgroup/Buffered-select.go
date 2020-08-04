
  
package main

import (
	"fmt"
	"time"
	"sync"
	
)
var goRoutine sync.WaitGroup
func one(ch chan string){
	defer goRoutine.Done()
	for i:=0;i<5;i++{
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
	for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch:
            fmt.Println("received", msg1)
        case msg2 := <-ch2:
            fmt.Println("received", msg2)
        }
    }
	goRoutine.Wait()
}
