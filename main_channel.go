package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	people := [4]string{"nico", "flynn", "dani", "haesoo"}
	for _, person := range people {
		go isSexy(person, channel)
	}
	// resultOne := <-channel
	// resultTwo := <-channel
	// fmt.Println("Waiting for messages")
	// fmt.Println("Received this message:", resultOne)
	// fmt.Println("Received this message:", resultTwo)
	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for ", i, " ")
		fmt.Println(<-channel)
	}
}

func isSexy(person string, channel chan string) {
	time.Sleep(time.Second * 5)
	channel <- person + " is sexy"
}
