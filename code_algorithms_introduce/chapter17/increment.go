package chapter17

func Increment(A *Calc) {
	i := 0
	for i < len(*A) && (*A)[i] == 1 {
		(*A)[i] = 0
		i += 1
	}
	if i < len(*A) {
		(*A)[i] = 1
	}
}

type Calc [32]int
