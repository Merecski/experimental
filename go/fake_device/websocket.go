package main

import (
	"log"
	"net"
	"net/http"

	"bitbucket.org/astiusa/discover-service/rpc"
	"github.com/gorilla/websocket"
)

// Interface for an application that supports clients connecting via Websocket.
type WebsocketApp interface {

	// Connect is called to create a new client of the app.
	Connect() WebsocketClient
}

// Interface for one client to interact with an app via messages.
type WebsocketClient interface {

	// SendChan returns a channel for sending messages to the client.
	SendChan() chan interface{}

	// ReceiveChan returns a channel for receiving messages from the client.
	ReceiveChan() chan interface{}

	// Disconnect the client from the app.
	Disconnect()
}

var upgrader = websocket.Upgrader{
	// Disable same origin check
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (wsi *WebsocketInterface) wsHandler(w http.ResponseWriter, r *http.Request) {

	// Upgrade HTTP -> Websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	doneChan := make(chan bool)

	// Create client interface to the application
	client := wsi.app.Connect()
	defer func() {
		// signal to goroutine that it's done
		doneChan <- true
		// disconnect the client
		client.Disconnect()
	}()
	var done bool

	// Pump responses from application to websocket
	go func() {
		for !done {
			select {
			case done = <-doneChan:
				break
			case msg := <-client.SendChan():
				ws.WriteJSON(msg)
			}
		}
		log.Println("Client disconnected")
	}()

	// Pump requests from websocket to application
	for {
		var req *rpc.Request
		if err := ws.ReadJSON(&req); err != nil {
			return
		}
		req.ResponseChan = client.SendChan()
		client.ReceiveChan() <- req
	}
}

// WebsocketInterface allows Websocket clients to interact with an Application
type WebsocketInterface struct {
	app      WebsocketApp
	listener net.Listener
}

// NewWebsocketInterface Creates a new WebsocketInterface to the provided app
func NewWebsocketInterface(app WebsocketApp, addr string) (*WebsocketInterface, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	wsi := &WebsocketInterface{app: app, listener: listener}
	mux := http.NewServeMux()
	mux.HandleFunc("/", wsi.wsHandler)
	go http.Serve(listener, mux)
	return wsi, nil
}

// Close the Websocket listener
func (wsi *WebsocketInterface) Close() {
	wsi.listener.Close()
	wsi.listener = nil
	// Any open websockets are not closed automatically.
	// Add logic to do that here if desired.
}
