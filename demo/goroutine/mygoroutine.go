package main

import (
	"fmt"
)

var c chan string

func Pingpong() {
	for {
		c<- fmt.Sprintf("Hi lllllllllly")
		fmt.Println(<-c)
	}
}

func main() {
	c = make(chan string)
	go Pingpong()
	for i := 0; i < 10; i++ {
			fmt.Println(<-c)
			c<- fmt.Sprintf("Hello lllllllllly")
	}
	
}