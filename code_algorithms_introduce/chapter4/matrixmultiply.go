package chapter4

// SquareMatrixMultiply 朴素的矩阵乘法，时间复杂度为O(n^3)
func SquareMatrixMultiply(A, B [][]int) [][]int {
	C := make([][]int, len(A))
	for i := range C {
		C[i] = make([]int, len(A[0]))
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			for k := 0; k < len(A); k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

// StrassenMatrixMultiply Strassen 方法
func StrassenMatrixMultiply(A, B [][]int) [][]int {
	n := len(A)
	C := createSquareSubMatrix(n)
	if n == 1 {
		C[0][0] = A[0][0] * B[0][0]
		return C
	} else {
		newSize := n / 2
		// 创建10个n/2 * n/2的矩阵S1,S2... S10
		S := make([][][]int, 11)
		for i := 0; i < 11; i++ {
			S[i] = createSquareSubMatrix(newSize)
		}
		// 得到子矩阵
		a11 := createSquareSubMatrix(newSize)
		a12 := createSquareSubMatrix(newSize)
		a21 := createSquareSubMatrix(newSize)
		a22 := createSquareSubMatrix(newSize)
		b11 := createSquareSubMatrix(newSize)
		b12 := createSquareSubMatrix(newSize)
		b21 := createSquareSubMatrix(newSize)
		b22 := createSquareSubMatrix(newSize)
		for i := 0; i < newSize; i++ {
			for j := 0; j < newSize; j++ {
				a11[i][j] = A[i][j]                 // top left
				a12[i][j] = A[i][j+newSize]         // top right
				a21[i][j] = A[i+newSize][j]         // bottom left
				a22[i][j] = A[i+newSize][j+newSize] // bottom right

				b11[i][j] = B[i][j]                 // top left
				b12[i][j] = B[i][j+newSize]         // top right
				b21[i][j] = B[i+newSize][j]         // bottom left
				b22[i][j] = B[i+newSize][j+newSize] // bottom right
			}
		}
		//S1=B12-B22
		subtract(b12, b22, S[1])
		//s2=A11+A12
		add(a11, a12, S[2])
		//s3=a21+a22
		add(a21, a22, S[3])
		//s4=b21-b11
		subtract(b21, b11, S[4])
		//s5=a11+a22
		add(a11, a22, S[5])
		//s6=b11+b22
		add(b11, b22, S[6])
		//s7=a12-a22
		subtract(a12, a22, S[7])
		//s8=b21+b22
		add(b21, b22, S[8])
		//s9=a11-a21
		subtract(a11, a21, S[9])
		//s10=b11+b12
		add(b11, b12, S[10])

		//P1=a11*s1
		P1 := StrassenMatrixMultiply(a11, S[1])
		//P2=S2*b22
		P2 := StrassenMatrixMultiply(S[2], b22)
		//P3=S3*b11
		P3 := StrassenMatrixMultiply(S[3], b11)
		//P4=a22*S4
		P4 := StrassenMatrixMultiply(a22, S[4])
		//P5=S5*S6
		P5 := StrassenMatrixMultiply(S[5], S[6])
		//P6=S7*S9
		P6 := StrassenMatrixMultiply(S[7], S[8])
		//P7=S9*S10
		P7 := StrassenMatrixMultiply(S[9], S[10])

		//创建结果矩阵
		c11 := createSquareSubMatrix(newSize)
		c12 := createSquareSubMatrix(newSize)
		c21 := createSquareSubMatrix(newSize)
		c22 := createSquareSubMatrix(newSize)

		//c11=P5+P4-P2+P6
		add(P5, P4, c11)
		subtract(c11, P2, c11)
		add(c11, P6, c11)

		//c12=P1+P2
		add(P1, P2, c12)

		//c21=P3+P4
		add(P3, P4, c21)

		//c22=P5+P1-P3-P7
		add(P5, P1, c22)
		subtract(c22, P3, c22)
		subtract(c22, P7, c22)

		//复制回去
		for i := 0; i < newSize; i++ {
			for j := 0; j < newSize; j++ {
				C[i][j] = c11[i][j]
				C[i][j+newSize] = c12[i][j]
				C[i+newSize][j] = c21[i][j]
				C[i+newSize][j+newSize] = c22[i][j]
			}
		}
	}

	return C
}

func createSquareSubMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return matrix
}

// subtract A-B to S
func subtract(A, B [][]int, S [][]int) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			S[i][j] = A[i][j] - B[i][j]
		}
	}
}
func add(A, B [][]int, S [][]int) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {

			S[i][j] = A[i][j] + B[i][j]
		}
	}
}
