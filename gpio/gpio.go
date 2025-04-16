package main

import (
	"machine"
	"time"
)

// benchmarkGPIOToggle toggles GPIO2 repeatedly and measures total time.
func benchmarkGPIOToggle() {
	const pin = machine.GPIO2
	const toggleCount = 1000

	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	println("task,method,toggles,total_time_us,avg_us")
	time.Sleep(3 * time.Second) // USB serial connect delay

	start := time.Now()

	for i := 0; i < toggleCount; i++ {
		pin.High()
		pin.Low()
	}

	duration := time.Since(start).Microseconds()
	avg := float64(duration) / float64(toggleCount)

	println("gpio,toggle,", toggleCount, ",", duration, ",", avg)
}
