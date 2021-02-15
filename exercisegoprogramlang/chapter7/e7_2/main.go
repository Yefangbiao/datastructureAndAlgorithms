package main

import (
	"bytes"
	"fmt"
	"io"
)

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var buf bytes.Buffer
	fmt.Fprint(&buf, w)
	c := new(int64)
	*c = int64(buf.Len())
	return &buf, c
}

func main() {
	var b bytes.Buffer
	b.Write([]byte("sadfihisadf"))
	w, _ := CountingWriter(&b)
	fmt.Println(w)
}
