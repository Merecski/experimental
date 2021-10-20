package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type emb1 struct {
	info int
}

type testStruct struct {
	data map[string]*emb1
}

var iter = 100000

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("hi?")
	x := &testStruct{
		data: map[string]*emb1{},
	}
	mtx := sync.Mutex{}

	start := make(chan struct{})
	finish := make(chan bool)

	go func() {
		<-start
		for i := 0; i < iter; i += 1 {
			tmp := &emb1{}
			tmp.info = rand.Int()
			log.Printf("write %d", i)
			mtx.Lock()
			x.data["a"] = tmp
			mtx.Unlock()
		}
		finish <- true
	}()
	go func() {
		<-start
		for i := 0; i < iter; i += 1 {
			mtx.Lock()
			tmp := x.data["a"]
			mtx.Unlock()
			log.Printf("Race? %d: %+v", i, tmp)
		}
		finish <- true
	}()

	log.Println("Lining up goroutines")
	time.Sleep(time.Second * 1)
	close(start)
	<-finish
	<-finish
}
