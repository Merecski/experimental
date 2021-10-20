package main

import (
	"fmt"
	"net"
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

func main() {
	temp := net.ParseIP("[fff1::ffff:2df2]")
	fmt.Println(net.ParseIP("192.0.2.1"))
	fmt.Println(temp)
	fmt.Println(net.ParseIP("192.0.2"))

}
