package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

func readTransformsFromFile(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	transforms := make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		transforms = append(transforms, sc.Text())
	}
	return transforms, nil
}

func main() {
	var fileOption = flag.String("file", "", "The file to read words from")
	flag.Parse()

	if *fileOption == "" {
		log.Fatal("--file option missing")
	}

	transforms, err := readTransformsFromFile(*fileOption)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}

}
