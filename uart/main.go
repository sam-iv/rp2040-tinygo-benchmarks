package main

import (
	"machine"
	"time"
)

func main() {
	machine.Serial.Configure(machine.UARTConfig{}) // USB serial
	time.Sleep(10 * time.Second)

	println("TinyGo UART Benchmark Starting...")
	benchmarkUART()

	for {
		time.Sleep(10 * time.Second)
	}
}
