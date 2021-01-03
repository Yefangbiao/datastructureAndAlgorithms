package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var SHA384 = flag.Bool("a", false, "output sha384")
var SHA512 = flag.Bool("b", false, "output sha512")

func main() {
	flag.Parse()
	// 第一个参数是输入的数据，第二个参数是是否指定其他输出方式
	var s string
	data := os.Args[1]
	switch {
	case *SHA384:
		s = fmt.Sprint(sha512.Sum384([]byte(data)))
	case *SHA512:
		s = fmt.Sprint(sha512.Sum512([]byte(data)))
	default:
		s = fmt.Sprint(sha256.Sum256([]byte(data)))
	}
	fmt.Printf("%x\n", s)
}
