package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func exists(domain string) (bool, error) {
	const whoIsServer string = "com.whois-servers.net"
	conn, err := net.Dial("tcp", whoIsServer+":43")
	if err != nil {
		return false, err
	}
	defer conn.Close()
	conn.Write([]byte(domain + "rn"))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		response := strings.ToLower(scanner.Text())
		if strings.Contains(response, "no match") {
			return false, nil
		}
	}
	return true, nil
}

var marks = map[bool]string{true: "✔", false: "✖"}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		domain := s.Text()
		fmt.Print(domain, " ")
		exist, err := exists(domain)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(marks[!exist])
		time.Sleep(1 * time.Second)
	}
}
