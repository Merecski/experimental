package main

import (
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"path/filepath"
	"time"

	"github.com/markbates/pkger"
	"github.com/webview/webview"
)

const (
	externalWebPrefix = "/dist"
)

type funtime struct {
	ID   int
	Name string
}

var f funtime

// Starts a server and serves all assets as needed
func startServer() string {
	ln, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer ln.Close()

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			path := filepath.Join(externalWebPrefix, r.URL.Path)

			if f, err := pkger.Open(path); err != nil {
				log.Println(err)
			} else {
				mimeType := mime.TypeByExtension(filepath.Ext(path))
				if filepath.Ext(path) == ".js" {
					mimeType = "text/javascript"
				}
				w.Header().Add("Content-Type", mimeType)
				io.Copy(w, f)
				f.Close()
			}
		})
		log.Fatal(http.Serve(ln, nil))
	}()

	log.Println("Webserver started at", ln.Addr().String())
	return "http://" + ln.Addr().String()
}

func printName(fun *funtime) {
	go func() {
		for {
			log.Println(fun.Name)
			time.Sleep(4 * time.Second)
		}
	}()
}

func returnData() (funtime, error) {
	return f, nil
}

func main() {
	baseURL := startServer()
	log.Println("BaseURL:", baseURL)
	f = funtime{
		ID:   47,
		Name: "test_webview",
	}

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	w.Bind("App", returnData)
	w.Navigate(baseURL + "/index.html")
	printName(&f)
	w.Run()
}
