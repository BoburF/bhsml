package basm

import (
	"os"

	astcontructor "ferxes.uz/bhsml/src/ast-contructor"
	"ferxes.uz/bhsml/src/parser"
	"ferxes.uz/bhsml/src/tokenizer"
)

func Assemble(input *os.File, output *os.File) error {
	tokenizer := tokenizer.NewTokenizer(input)
	parser := parser.NewParser(*tokenizer)
	parser.Parse()
	ast := astcontructor.NewASTConstructor(parser.Stack)
	ast.Construct()

	json, err := ast.ToJSON()
	if err != nil {
		return err
	}

	if _, err := output.Write([]byte(json)); err != nil {
		return err
	}

	return nil
}
