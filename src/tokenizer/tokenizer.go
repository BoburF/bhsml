package tokenizer

import (
	"bufio"
	"io"
)

type Tokenizer struct {
	Tokenize *bufio.Reader
	Index    int
	Line     int
}

func NewTokenizer(reader io.Reader) *Tokenizer {
	tokenizer := &Tokenizer{
		Tokenize: bufio.NewReader(reader),
		Index:    0,
		Line:     1,
	}

	return tokenizer
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
	if b == 10 {
		tk.Line++
	}

	tk.Index++

	return b, nil
}

func (tk *Tokenizer) PeekToNext() (byte, error) {
	b, err := tk.Tokenize.ReadByte()
	if err != nil {
		return 0, err
	}

    tk.Tokenize.UnreadByte()

    return b, nil
}
