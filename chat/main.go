package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"trace"
)

func main() {

	var addr = flag.String("addr", ":8080", "The addr of the application")
	flag.Parse()

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// get the room going
	go r.run()

	// start the web server
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
