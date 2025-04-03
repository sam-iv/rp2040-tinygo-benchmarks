package main

import (
	"machine"
	"time"
)

func benchmarkBubbleSort() {
	sizes := []int{10, 50, 100}
	println("task,method,size,time_us")

	for _, size := range sizes {
		// Create reversed array
		data := make([]int, size)
		for i := 0; i < size; i++ {
			data[i] = size - i
		}

		start := time.Now()
		BubbleSort(data)
		elapsed := time.Since(start).Microseconds()

		println("bubblesort,bubble,", size, ",", elapsed)
	}
}

func main() {
	machine.Serial.Configure(machine.UARTConfig{})
	time.Sleep(time.Second * 10) // allow serial connection

	println("🔧 TinyGo Bubble Sort Benchmark Starting...")
	benchmarkBubbleSort()

	for {
		time.Sleep(time.Second * 10)
	}
}
