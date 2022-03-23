package main

import (
	"fmt"
	"strconv"
	"time"
)

var testmap map[string]bool
var listLookup []string
var iterations int

func setup() {
	iterations = 1000
	testmap = make(map[string]bool)
	for i := 0; i < iterations; i++ {
		id := strconv.Itoa(i)
		testmap[id] = true
		listLookup = append(listLookup, id)
	}
}

func run1() {
	for index, value := range listLookup {
		if _, ok := testmap[value]; !ok {
			fmt.Println("ERROR")
		}
		if index == 400 {
			return
		}
		if testmap[value] {
			return
		}
	}
}

func run2() {
	for key, value := range testmap {
		if !value {
			return
		} else {
			key += "_done"
		}
	}
}

func run3() {
	for index, value := range listLookup {
		if index == 100 {
			return
		}
		if value == "400" {
			return
		}
	}
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func main() {
	setup()
	fmt.Printf("Running loops for %d iterations\n", iterations)
	func() {
		defer elapsed("run 1 (slice w/ map)")()
		run1()
	}()
	func() {
		defer elapsed("run 2 (map only)")()
		run2()
	}()
	func() {
		defer elapsed("run 3 (slice only)")()
		run3()
	}()

}
