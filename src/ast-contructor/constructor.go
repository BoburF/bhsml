package astcontructor

import (
	"encoding/json"

	"ferxes.uz/bhsml/src/parser"
)

type ASTContructor struct {
	Tokens []parser.Tag
	Tree   *Node
}

func NewASTConstructor(tokens []parser.Tag) *ASTContructor {
	return &ASTContructor{
		Tokens: tokens,
	}
}

func (ast *ASTContructor) Construct() {
	ast.Tree = &Node{Type: "bhsml", Children: make([]*Node, 0)}
	stack := []*Node{ast.Tree}

	for _, token := range ast.Tokens {
		if token.Name[0] == '/' {
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
		} else {
			newNode := Node{
				Type:     token.Type,
				Attrs:    token.Atributes,
				Text:     token.Name,
				Children: make([]*Node, 0),
			}

			if newNode.Type == "text" {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, &newNode)
			} else {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, &newNode)
				stack = append(stack, &newNode)
			}
		}
	}
}

func (ast *ASTContructor) ToJSON() (string, error) {
	jsonData, err := json.MarshalIndent(ast.Tree, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
