package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, err := readInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	words := strings.Fields(text)
	fmt.Println(len(words))
}

func readInput() (src string, err error) {
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	if src == "" {
		return src, errors.New("empty string")
	}
	return src, nil
}
