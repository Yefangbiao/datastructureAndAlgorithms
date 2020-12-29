package p1_3

import (
	"strconv"
	"strings"
)

func echo() {
	var s, sep string
	arr := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = strconv.Itoa(i)
	}
	for i := 0; i < 1000; i++ {
		s += arr[i] + sep
		sep = " "
	}
}

func echoJoins() {
	var sep string
	arr := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = strconv.Itoa(i)
	}
	_ = strings.Join(arr, sep)
}
