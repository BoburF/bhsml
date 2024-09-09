package main

import (
	"fmt"
	"os"

	basm "ferxes.uz/bhsml/src/bhsml-asm"
)

func main() {
	file, err := os.Open("example.bhsml")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fo, err := os.Create("ast.json")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	err = basm.Assemble(file, fo)
	fmt.Println("JSON data written to ast.json successfully")
}
