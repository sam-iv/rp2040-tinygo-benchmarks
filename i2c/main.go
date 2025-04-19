package main

import (
	"machine"
	"time"
)

func main() {
	machine.Serial.Configure(machine.UARTConfig{})
	time.Sleep(10 * time.Second) // Wait for USB serial to connect

	println("TinyGo I2C Benchmark Starting...")
	benchmarkI2C()

	for {
		time.Sleep(time.Second * 10)
	}
}
