package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, link := range links {
		//checkLink(link)
		go checkLink(link, c)
	}
	//	fmt.Println(<-c) //blocking call
	//fmt.Println(<-c)  dead lock
	//for i := 0; i < len(links); i++ {
	for l := range c {
		//fmt.Println(<-c)
		//time.Sleep(1 * time.Second)  not appropriate to put sleep in main routine
		go func(link string) {
			time.Sleep(1 * time.Second)
			checkLink(link, c)
		}(l)
	}
}
func checkLink(link string, c chan string) {
	//time.Sleep(2 * time.Second)
	_, err := http.Get(link) //blocking code
	if err != nil {
		fmt.Println(link + " might be down")
		//	c <- "might be down"
		c <- link
		return
	}
	c <- link
	fmt.Println(link + " working fine")
}
