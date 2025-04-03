package main

import (
	"machine"
	"time"
)

func benchmarkQuickSort() {
	sizes := []int{10, 50, 100}
	println("task,method,size,time_us")

	for _, size := range sizes {
		// Reversed array = worst case
		data := make([]int, size)
		for i := 0; i < size; i++ {
			data[i] = size - i
		}

		start := time.Now()
		QuickSort(data, 0, size-1)
		elapsed := time.Since(start).Microseconds()

		println("quicksort,quick,", size, ",", elapsed)
	}
}

func main() {
	machine.Serial.Configure(machine.UARTConfig{})
	time.Sleep(time.Second * 10)

	println("🔧 TinyGo Quick Sort Benchmark Starting...")
	benchmarkQuickSort()

	for {
		time.Sleep(time.Second * 10)
	}
}
