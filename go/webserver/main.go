package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

type Server struct {
	close chan bool
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s", data)
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Shutting down...")
}

func crypto(w http.ResponseWriter, r *http.Request) {
	cryp := pat.Param(r, "data")
	s := struct {
		Name string
	}{
		cryp,
	}
	out, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func frontendListener() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/crypto/:data"), crypto)
	http.ListenAndServe("localhost:8047", mux)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	mux.HandleFunc(pat.Get("/shutdown"), shutdown)
	mux.HandleFunc(pat.Get("/"), serveHome)
	fmt.Println("Starting HTTP Listeners")
	go frontendListener()
	http.ListenAndServe("localhost:8000", mux)
}
