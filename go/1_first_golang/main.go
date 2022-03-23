package main

import (
	"fmt"
	"time"
)

type st1 struct {
	value string
}

type st2 struct {
	value string
}

func foo1(s *st1) {
	s.value = "hello9"
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func HelloWorld() {
	msg := Hello("John")
	fmt.Printf("%s\n", msg)
	x := 5 - 10
	fmt.Println("Variable x contains ", x, " as it's value.")
}

func optionalArgs(arg1 string, arg2 string) {
	var mainArg string
	if arg1 == "" {
		mainArg = arg2
	} else {
		mainArg = arg1
	}
	fmt.Println(mainArg)
}

func foo1_runner() {
	tmp := st1{}
	foo1(&tmp)
	fmt.Println(tmp)
}

func test1(kill <-chan bool) {
	c1 := make(chan string)
	go func() {
		msg := "$$$$"
		for {
			select {
			case c1 <- msg:
			case <-kill:
				return
			default:
				fmt.Printf("hehe overflow\n")
				time.Sleep(time.Millisecond * 1)
			}
		}
	}()
	go func() {
		msg := "haha"
		for {
			select {
			case c1 <- msg:
			case <-kill:
				return
			default:
				fmt.Printf("haha overflow\n")
				time.Sleep(time.Millisecond * 1)
			}
		}
	}()

	for {
		select {
		case rec := <-c1:
			fmt.Println(rec)
		case <-kill:
			return
		}
	}
}

func test2(kill <-chan bool) {
	c1 := make(chan string)
	go func() {
		for {
			select {
			case msg := <-c1:
				fmt.Printf("Routine 1 got msg: %s\n", msg)
			case <-kill:
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case msg := <-c1:
				fmt.Printf("Routine 2 got msg: %s\n", msg)
			case <-kill:
				return
			}
		}
	}()

	msg := "o7"
	for {
		select {
		case c1 <- msg:
		case <-kill:
			return
		}
		break
	}
}

func main() {
	kill := make(chan bool)
	time.Sleep(time.Second * 1)
	go test2(kill)
	time.Sleep(time.Second * 1)
	kill <- true
}
