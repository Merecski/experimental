package main

import (
	"encoding/json"
	"log"
	"net"
	"time"

	"bitbucket.org/astiusa/deltaapi"
)

type FakeDevice struct {
	txConn          *net.UDPConn
	lastTime        time.Time
	name            string
	macAddress      string
	firmwareVersion string
	ipAddress       string
}

func NewFakeDevice(name, txAddrStr, macAddress, firmwareVersion, ipAddress string) (*FakeDevice, error) {

	txAddr, err := net.ResolveUDPAddr("udp", txAddrStr)
	if err != nil {
		return nil, err
	}
	txConn, err := net.DialUDP("udp", nil, txAddr)
	if err != nil {
		return nil, err
	}

	return &FakeDevice{txConn: txConn, name: name, macAddress: macAddress, firmwareVersion: firmwareVersion, ipAddress: ipAddress}, nil
}

func (fd *FakeDevice) Poll() {
	now := time.Now()
	dt := now.Sub(fd.lastTime).Seconds()
	if dt > 5.0 {
		log.Println("Sending existence", fd.name)
		buf, _ := json.Marshal(deltaapi.Existence{
			PacketType: "existence",
			Device:     fd.name,
			DeviceID:   fd.macAddress,
			Version:    fd.firmwareVersion,
			IPAddr:     fd.ipAddress,
		})
		fd.txConn.Write(buf)
		fd.lastTime = now
	}
}
