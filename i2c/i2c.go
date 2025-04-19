package main

import (
	"machine"
	"time"
)

// benchmarkI2C measures write latency over I2C0 to a passive slave device.
//
// Sends 100 identical 4-byte messages to address 0x42 over I2C0.
// Measures total and average time taken. Results printed in CSV format.
//
// SDA: GPIO8 (pin 11), SCL: GPIO9 (pin 12)
// Output:
//   task,method,iterations,total_time_us,avg_time_us
func benchmarkI2C() {
	const addr = 0x42
	const loops = 100
	data := []byte{'T', 'e', 's', 't'}

	machine.I2C0.Configure(machine.I2CConfig{
		SCL: machine.GPIO9,
		SDA: machine.GPIO8,
		Frequency: 100000, // 100 kHz
	})

	println("Benchmark: I2C Write")
	println("task,method,iterations,total_time_us,avg_time_us")

	start := time.Now()

	for i := 0; i < loops; i++ {
		machine.I2C0.Tx(uint16(addr), data, nil)
	}

	elapsed := time.Since(start).Microseconds()
	avg := float64(elapsed) / float64(loops)

	println("i2c,write,", loops, ",", elapsed, ",", avg)
}
