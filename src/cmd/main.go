package main

import (
	"fmt"
	"os"

	"ferxes.uz/bhsml/src/parser"
	"ferxes.uz/bhsml/src/tokenizer"
)

func main() {
	file, err := os.Open("example.bhsml")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	tokenizer := tokenizer.NewTokenizer(file)
    parser := parser.NewParser(*tokenizer)
    parser.Parse()
    fmt.Println(parser.Stack)
}
