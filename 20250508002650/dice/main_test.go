package main

import "testing"

func BenchmarkDiceRollFunctional(b *testing.B) {
	for b.Loop() {
		createHistogram([]int{5, 1000, 100, 30, 5})
	}
}

func BenchmarkDiceRollRecursive(b *testing.B) {
	for b.Loop() {
		createHistogramRecursion([]int{5, 1000, 100, 30, 5})
	}
}
