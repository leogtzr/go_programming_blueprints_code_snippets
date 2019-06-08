package main

import (
	"fmt"
	"log"

	"github.com/joeshaw/envdecode"
)

func main() {
	var ts struct {
		Home  string `env:"HOME, required"`
		Shell string `env:"SHELL, required"`
	}

	if err := envdecode.Decode(&ts); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(ts)
}
