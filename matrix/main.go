package main

import (
	"machine"
	"time"
)

func createMatrix(n int, fill func(i, j int) int) [][]int {
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
		for j := range mat[i] {
			mat[i][j] = fill(i, j)
		}
	}
	return mat
}

func benchmarkMatrixMult() {
	sizes := []int{10, 20}
	println("task,method,size,time_us")

	for _, size := range sizes {
		A := createMatrix(size, func(i, j int) int { return i + j })
		B := createMatrix(size, func(i, j int) int { return i - j })
		C := createMatrix(size, func(i, j int) int { return 0 })

		start := time.Now()
		MatrixMultiply(size, A, B, C)
		elapsed := time.Since(start).Microseconds()

		println("matrix,multiply,", size, ",", elapsed)
	}
}

func main() {
	machine.Serial.Configure(machine.UARTConfig{})
	time.Sleep(time.Second * 10)

	println("🔧 TinyGo Matrix Mult Benchmark Starting...")
	benchmarkMatrixMult()

	for {
		time.Sleep(time.Second * 10)
	}
}
