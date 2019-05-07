package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"trace"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
)

func main() {

	var addr = flag.String("addr", ":8080", "The addr of the application")
	flag.Parse()

	gomniauth.SetSecurityKey("a6da4e35e6b0cb614e93e33052dcf9596f21d23c")
	gomniauth.WithProviders(
		//facebook.New("key", "secret", "http://localhost:8080/auth/callback/facebook"),
		github.New("81a518c207b48f360280", "a6da4e35e6b0cb614e93e33052dcf9596f21d23c", "http://localhost:8080/auth/callback/github"),
		//google.New("key", "secret", "http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	// get the room going
	go r.run()

	// start the web server
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}