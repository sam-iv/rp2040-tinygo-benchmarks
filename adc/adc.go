package main

import (
	"machine"
	"time"
)

// benchmarkADC measures ADC read latency on GPIO26 (ADC0).
//
// It reads 1000 samples, tracks total time, and prints the average read latency.
// Output is printed using println() only — no fmt or strconv.
//
// Output format:
//   task,method,reads,avg_time_us
func benchmarkADC() {
	const samples = 1000

	machine.InitADC()

	adc := machine.ADC{Pin: machine.ADC0}
	adc.Configure(machine.ADCConfig{})

	println("task,method,reads,avg_time_us")

	var totalTime int64

	for i := 0; i < samples; i++ {
		start := time.Now()
		val := adc.Get()
		elapsed := time.Since(start).Microseconds()
		totalTime += elapsed

		if i%250 == 0 {
			println("adc,sample,", i, ",", val)
		}
	}

	avg := totalTime / int64(samples)
	println("adc,single_read,", samples, ",", avg)
}

// TODO: Comment