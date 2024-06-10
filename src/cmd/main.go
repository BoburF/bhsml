package main

import (
	"fmt"
	"io"
	"os"

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

	for {
		ch, err := tokenizer.Next()

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error:", err)
			break
		}
		fmt.Println(ch)
		fmt.Printf("Character: %c\n", ch)
	}
}
