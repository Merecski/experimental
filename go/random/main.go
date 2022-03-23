package main

import (
	"fmt"
	"sync"
)

var lok = sync.RWMutex{}
var glist = []*obj{}
var newList = []*obj{}
var start = make(chan bool, 2)
var finished = make(chan bool)

type obj struct {
	id int
}

func foo() {
	for i := 0; i < 1000000; i++ {
		tmp := &obj{
			id: i,
		}
		glist = append(glist, tmp)
	}
	go func() {
		search := 444777
		<-start
		lok.RLock()
		fmt.Println("Starting first routine!")
		for _, o := range glist {
			if o.id == search {
				lok.Lock()
				fmt.Println("Found it!")
				o.id = -1
				lok.Unlock()
				fmt.Println("Wrote to it!")
			}
		}
		lok.RUnlock()
		finished <- true
	}()
	start <- true
	lok.RLock()
	fmt.Println("Starting second routine!")
	for _, o := range glist {
		if o != nil {
			newList = append(newList, o)
		}
	}
	lok.RUnlock()
	<-finished
}

func main() {
	foo()
}
