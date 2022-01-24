package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: wordcount PHRASE")
		os.Exit(1)
	}
	phrase := os.Args[1]
	words := strings.Fields(phrase)
	fmt.Println(len(words))
}
