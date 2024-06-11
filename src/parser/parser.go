package parser

import (
	"bytes"
	"io"
	"strings"

	"ferxes.uz/bhsml/src/tokenizer"
)

type Parser struct {
	Tokenizer tokenizer.Tokenizer
	Stack     []Tag
	tempTag   *Tag
}

func NewParser(tokenizer tokenizer.Tokenizer) *Parser {
	return &Parser{
		Tokenizer: tokenizer,
	}
}

func (p *Parser) Parse() {
	byte, err := p.Tokenizer.Peek()
	if err != nil {
		if err == io.EOF {
			return
		}
	}
	if byte == 10 || byte == 32 {
		p.Tokenizer.Next()
		p.Parse()
		return
	}

	switch string(byte) {
	case "<":
		p.tempTag = &Tag{}
		p.tempTag.Position.StartIndex = p.Tokenizer.Index
		p.tempTag.Line = p.Tokenizer.Line

		p.Tokenizer.Next()
		tag, err := p.readUntil('>')
		if err != nil {
			return
		}
		p.parseAtributes(tag)

		p.tempTag.Type = "tag"

		p.Parse()
		break
	case ">":
		p.tempTag.Position.EndIndex = p.Tokenizer.Index
		p.Stack = append(p.Stack, *p.tempTag)

		p.Tokenizer.Next()
		p.Parse()
		break
	default:
		p.tempTag = &Tag{}
		p.tempTag.Position.StartIndex = p.Tokenizer.Index
		p.tempTag.Line = p.Tokenizer.Line

		text, err := p.readUntil('<')
		if err != nil {
			return
		}

		p.tempTag.Type = "text"
		p.tempTag.Name = text
		p.tempTag.Position.EndIndex = p.Tokenizer.Index
		p.Stack = append(p.Stack, *p.tempTag)

		p.Parse()
		return
	}

	return
}

func (p *Parser) readUntil(delimiter byte) (string, error) {
	var buf bytes.Buffer

	for {
		b, err := p.Tokenizer.PeekToNext()
		if err != nil {
			return "", err
		}
		if b == delimiter {
			break
		}
		buf.WriteByte(b)
		p.Tokenizer.Next()
	}

	return buf.String(), nil
}

func (p *Parser) parseAtributes(tag string) {
	words := strings.Fields(tag)

	if len(words) > 0 {
		p.tempTag.Name = words[0]
	}
	if len(words) <= 1 {
		return
	}

	p.tempTag.Atributes = make(map[string]string)
	for _, word := range words[1:] {
		parts := strings.SplitN(word, "=", 2)

		if len(parts) >= 2 {
			p.tempTag.Atributes[parts[0]] = parts[1]
		} else {
			p.tempTag.Atributes[parts[0]] = ""
		}
	}
}
