package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type S struct {
	S   string
	len int
}

func (s *S) Read(p []byte) (n int, err error) {
	if s.len >= len(s.S) {
		return 0, io.EOF
	}
	l := copy(p, s.S[s.len:])
	s.len = l
	return l, nil
}

func NewReader(s string) io.Reader {
	return &S{s, 0}
}

func main() {
	r := NewReader("123")
	s, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(s)
	}
	fmt.Println(string(s))
}
