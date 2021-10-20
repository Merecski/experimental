package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/pkg/browser"
)

type Errors struct {
	id  int
	msg string
}

func testStructs() {
	errSlice := []*Errors{}

	err1 := &Errors{
		id:  47,
		msg: "hi",
	}

	errSlice = append(errSlice, err1)

	err2 := &Errors{
		id:  47,
		msg: "hi",
	}

	for _, err := range errSlice {
		if err == err2 {
			fmt.Println(true)
		}
		fmt.Println(*err, *err2)
	}
}

type data struct {
	id   int
	info string
}

type testData1 struct {
	MapType    map[string]bool    `json:"map_type"`
	StructType map[int]*reqStates `json:"req_state"`
}

type reqStates struct {
	Submitted bool `json:"submitted"`
	Getting   bool `json:"getting"`
	Resetting bool `json:"resetting"`
}

func foo4() {
	x := &testData1{}
	x.StructType = make(map[int]*reqStates)
	// x.StructType[0] = &reqStates{}
	y := make(map[string]interface{})
	y["getting"] = true
	jsonString, _ := json.Marshal(y)
	fmt.Println(string(jsonString))
	err := json.Unmarshal(jsonString, x.StructType[0])
	fmt.Println(x.StructType[0])
	fmt.Println(err)
}

func openBrowser() {
	browser.OpenURL("http://www.google.com")
}

func netTest() {
	localInterfaces := []net.IP{}
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip.To4() != nil {
				localInterfaces = append(localInterfaces, ip)
			}
		}
	}

	fmt.Printf("Avaliable Networks: %v\n\n", localInterfaces)

	A := "172.17.0.0/32"
	B := "172.17.1.2/32"

	ipA, ipnetA, _ := net.ParseCIDR(A)
	ipB, ipnetB, _ := net.ParseCIDR(B)

	fmt.Println("Network address A: ", A)
	fmt.Println("IP address      B: ", B)
	fmt.Println("ipA              : ", ipA)
	fmt.Println("ipnetA           : ", ipnetA)
	fmt.Println("ipB              : ", ipB)
	fmt.Println("ipnetB           : ", ipnetB)

	fmt.Printf("\nDoes A (%s) contain: B (%s)?\n", ipnetA, ipB)

	if ipnetA.Contains(ipB) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func main() {
	netTest()
}
