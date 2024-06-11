package main

import (
	"fmt"
	"os"

	astcontructor "ferxes.uz/bhsml/src/ast-contructor"
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
	ast := astcontructor.NewASTConstructor(parser.Stack)
	ast.Construct()

	json, err := ast.ToJSON()
	if err != nil {
		fmt.Println("Error parsing to json:", err)
		return
	}

	fo, err := os.Create("ast.json")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := fo.Write([]byte(json)); err != nil {
		panic(err)
	}

	fmt.Println("JSON data written to ast.json successfully")
}
