package main

import (
	"fmt"
	"time"
)

var conn chan string
var kill chan bool

func foo1(name string) {
	for msg := range conn {
		fmt.Println(name, "got", msg)
	}
}

func runFoo1() {
	conn = make(chan string)
	kill = make(chan bool)

	go foo1("f1")
	go foo1("f2")

	time.Sleep(10 * time.Microsecond)
	for i := 0; i < 5; i++ {
		conn <- fmt.Sprintf("%d", i)
	}
	time.Sleep(1000 * time.Microsecond)
	close(conn)
}

func foo2() {
	ch := make(chan string, 10)
	ch <- "hi"
	go func() {
		for {
			out, ok := <-ch
			if !ok {
				fmt.Println("This is not ok")
				break
			}
			fmt.Println(out)
		}
		fmt.Println("exiting")
	}()
	time.Sleep(1 * time.Second)
	close(ch)
	fmt.Println("Closed")
	time.Sleep(1 * time.Second)
}

func main() {
	foo2()
}
