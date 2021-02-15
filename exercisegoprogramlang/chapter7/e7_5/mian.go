package main

import "io"

type S struct {
	r io.Reader
	N int64
}

func (s *S) Read(p []byte) (n int, err error) {
	if s.N < 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > s.N {
		p = p[0:s.N]
	}
	n, err = s.r.Read(p)
	s.N -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &S{r, n}
}
