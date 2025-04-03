package main

import (
	"machine"
	"math"
	"time"
)

func benchmarkFFT() {
	const N = 128
	real := make([]float32, N)
	imag := make([]float32, N)

	// Fill with a sine wave
	for i := 0; i < N; i++ {
		real[i] = float32(math.Sin(2 * math.Pi * float64(i) / float64(N)))
		imag[i] = 0
	}

	start := time.Now()
	fftRadix2(real, imag)
	elapsed := time.Since(start).Microseconds()

	println("task,method,size,time_us")
	println("fft,radix2,", N, ",", elapsed)
}

func main() {
	machine.Serial.Configure(machine.UARTConfig{})
	time.Sleep(time.Second * 10)

	println("🔧 TinyGo FFT Benchmark Starting...")
	benchmarkFFT()

	for {
		time.Sleep(time.Second * 10)
	}
}
