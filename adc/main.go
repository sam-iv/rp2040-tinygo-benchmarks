package main

import (
	"machine"
	"time"
)

func main() {
	machine.Serial.Configure(machine.UARTConfig{})
	time.Sleep(10 * time.Second)

	println("TinyGo ADC Benchmark Starting...")
	benchmarkADC()

	for {
		time.Sleep(10 * time.Second)
	}
}
// TODO: Comment