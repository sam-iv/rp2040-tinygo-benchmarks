package main

import (
	"machine"
	"time"
)

func main() {
	machine.Serial.Configure(machine.UARTConfig{})
	time.Sleep(10 * time.Second)

	println("TinyGo GPIO Toggle Benchmark Starting...")
	benchmarkGPIOToggle()

	for {
		time.Sleep(10 * time.Second)
	}
}
