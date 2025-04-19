package main

import (
	"machine"
	"time"
)

const (
	buttonPin = machine.GPIO14
	buzzerPin = machine.GPIO15
)

var (
	triggered    = false
	irqStart     int64
	irqLatency   int64
)

func handleInterrupt(pin machine.Pin) {
	irqLatency = time.Since(time.UnixMicro(irqStart)).Microseconds()
	println("interrupt,triggered,", irqLatency)
	triggered = true
}

// benchmarkInterrupt sets up GPIO interrupt on GPIO14 and tracks ISR latency.
func benchmarkInterrupt() {
	println("Benchmark: Interrupt Latency")
	println("task,method,latency_us")
	println("Press the button to trigger interrupt...")

	buttonPin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	buzzerPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	buzzerPin.Low()

	buttonPin.SetInterrupt(machine.PinToggle, handleInterrupt)

	for {
		irqStart = time.Now().UnixMicro()

		if triggered {
			buzzerPin.High()
			time.Sleep(50 * time.Millisecond)
			buzzerPin.Low()
			triggered = false
		}

		time.Sleep(1 * time.Millisecond)
	}
}
