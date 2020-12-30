package main

import "fmt"

// 以B为单位
const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	// although can not display ZB and YB, but it exist
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB/GB, YB/GB)
}
