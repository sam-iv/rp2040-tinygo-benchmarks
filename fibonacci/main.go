package main

import (
	"machine"
	"time"
)

func benchmarkFibonacci() {
	ns := []int{10, 20, 30, 35}

	println("task,method,n,result,time_us")

	for _, n := range ns {
		// Iterative
		start := time.Now()
		result := FibIterative(n)
		elapsed := time.Since(start).Microseconds()
		println("fibonacci,iterative,", n, ",", result, ",", elapsed)

		// Recursive
		start = time.Now()
		result = FibRecursive(n)
		elapsed = time.Since(start).Microseconds()
		println("fibonacci,recursive,", n, ",", result, ",", elapsed)
	}
}

func main() {
	// Init USB serial output
	machine.Serial.Configure(machine.UARTConfig{})
	time.Sleep(time.Second * 10) // wait for host to connect

	println("🔧 TinyGo Benchmark Runner Starting...")
	benchmarkFibonacci()

	for {
		time.Sleep(time.Second * 5)
	}
}
