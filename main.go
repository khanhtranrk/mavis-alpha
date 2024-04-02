package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveMaster(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "resource/html/master.html")
}

func serveAsis(w http.ResponseWriter, r *http.Request) {
}

func main() {
	flag.Parse()
	hub := newHub()
  broker := NewBroker(hub)
	go hub.run()
  go broker.Listen()
	http.HandleFunc("/", serveMaster)
	http.HandleFunc("/ds", serveAsis)
	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
