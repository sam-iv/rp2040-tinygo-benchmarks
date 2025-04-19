package main

import (
	"machine"
	"time"
)

// benchmarkUART sends a fixed message 100 times over UART0.
//
// Measures total time and calculates average time per message.
//
// Output format:
//   task,method,iterations,total_time_us,avg_time_us
func benchmarkUART() {
	const iterations = 100
	const baud = 115200
	const txPin = machine.GPIO0
	message := []byte("Hello from main Pico\n")

	uart := machine.UART0
	uart.Configure(machine.UARTConfig{
		TX:   txPin,
		BaudRate: baud,
	})

	println("Benchmark: UART Send")
	println("task,method,iterations,total_time_us,avg_time_us")

	start := time.Now()

	for i := 0; i < iterations; i++ {
		for _, b := range message {
			uart.WriteByte(b)
		}
	}

	duration := time.Since(start).Microseconds()
	avg := float64(duration) / float64(iterations)

	println("uart,send,", iterations, ",", duration, ",", avg)
}
