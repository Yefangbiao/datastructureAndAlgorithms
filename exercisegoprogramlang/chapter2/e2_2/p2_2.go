package p2_2

import (
	"fmt"
	"os"
	"strconv"
)

type kilogram float64

func main() {
	// 将数字转换为千克表示
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
		}
		fmt.Println(FToK(t))
	}
}

func FToK(num float64) kilogram {
	return kilogram(num)
}

func (k kilogram) String() string {
	s := fmt.Sprintf("%g kilogram\n", k)
	return s
}
