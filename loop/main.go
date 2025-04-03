package main

import (
	"machine"
	"time"
)

func benchmarkLoopOverhead() {
	iterations := []int{1000, 10000, 100000, 1000000}
	println("task,method,iterations,time_us")

	for _, n := range iterations {
		start := time.Now()
		result := LoopN(n) // ✅ use the result to prevent optimization
		elapsed := time.Since(start).Microseconds()

		// Optionally include result in output to guarantee use
		println("loop,for_loop,", n, ",", elapsed, ",", "result:", result)
	}
}

func main() {
	machine.Serial.Configure(machine.UARTConfig{})
	time.Sleep(time.Second * 10) // allow serial monitor to connect

	println("🔧 TinyGo Loop Overhead Benchmark Starting...")
	benchmarkLoopOverhead()

	for {
		time.Sleep(time.Second * 10)
	}
}
