package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bitbucket.org/astiusa/deltaapi"
)

var addr = flag.String("addr", "10.2.101.26:8085", "http service address")
var fakeDevName = "TESTING_DEVICE_1234"

func main() {
	var err error
	var fake *FakeDevice
	if fakeDevName != "" {
		txAddr := fmt.Sprintf("%s:%d", deltaapi.DeviceMcastAddress, deltaapi.DeviceExistencePort)
		fake, err = NewFakeDevice(fakeDevName, txAddr, "a1:2b:3c:4d:5e:63", "v0.9.0", "10.2.101.121/16")
		if err != nil {
			log.Fatalln("Failed to create fake device:", err)
		}
	}

	doneChan := make(chan bool, 1)
	var done bool

	// Cleanup on SINGINT, SIGTERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-c
		log.Printf("Signal %v recevied", sig)
		doneChan <- true
	}()

	for !done {
		select {
		case done = <-doneChan:
			break
		default:
		}
		if done {
			break
		}
		if fake != nil {
			fake.Poll()
		}
		time.Sleep(time.Millisecond * 10)
	}
}
