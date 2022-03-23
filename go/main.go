package main

import "fmt"

func killYourself(with interface{}, done chan<- bool) {
	defer func() { done <- true }()
	switch with.(type) {
	case (string):
		fmt.Printf("You have been ended with a %s\n", with)
	default:
		fmt.Println("You commit stop life")
	}
}

func main() {
	done := make(chan bool)

	go killYourself("rusty spoon", done)

	<-done
}
