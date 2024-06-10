package tokenizer

import (
	"bufio"
	"io"
)

type Tokenizer struct {
	Tokenize *bufio.Reader
	index    int
}

func NewTokenizer(reader io.Reader) *Tokenizer {
	return &Tokenizer{
		Tokenize: bufio.NewReader(reader),
		index:    0,
	}
}

func (tk *Tokenizer) Peek() (byte, error) {
	bs, err := tk.Tokenize.Peek(1)
	if err != nil {
		return 0, err
	}

	return bs[0], nil
}

func (tk *Tokenizer) Next() (byte, error) {
	b, err := tk.Tokenize.ReadByte()
	if err != nil {
		return 0, err
	}

	tk.index++

	return b, nil
}
