package main

var LoopSink int

func LoopN(n int) int {
	for i := 0; i < n; i++ {
		LoopSink++
	}
	return LoopSink
}
