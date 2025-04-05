// benchmarkLoopOverhead runs a benchmark to measure basic loop iteration overhead.
//
// This benchmark uses a global counter and executes for-loops with a fixed number
// of iterations (1K to 1M). Results are printed in CSV format to evaluate the
// raw overhead of loop control logic in TinyGo.
//
// Output format:
//   task,method,iterations,time_us[,result]
func benchmarkLoopOverhead() {
	iterations := []int{1000, 10000, 100000, 1000000}
	println("task,method,iterations,time_us")

	for _, n := range iterations {
		start := time.Now()
		result := LoopN(n) // use the result to prevent optimization
		elapsed := time.Since(start).Microseconds()

		// Optionally include result in output to verify loop execution
		println("loop,for_loop,", n, ",", elapsed, ",", "result:", result)
	}
}

func main() {
	machine.Serial.Configure(machine.UARTConfig{})
	time.Sleep(time.Second * 10) // allow serial monitor to connect

	println("TinyGo Loop Overhead Benchmark Starting...")
	benchmarkLoopOverhead()

	for {
		time.Sleep(time.Second * 10)
	}
}
