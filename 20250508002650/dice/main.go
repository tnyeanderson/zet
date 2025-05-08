package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Array of dice, where the value is the number of sides
	dice := []int{}
	maxVal := 0
	for _, arg := range os.Args[1:] {
		sides, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("not a number: %s\n", arg)
		}
		dice = append(dice, sides)
		maxVal += sides
	}

	// Create histograms
	histogram := createHistogram(dice)
	histogramR := createHistogramRecursion(dice)

	printHistogram(histogram, maxVal)
	printHistogram(histogramR, maxVal)
}

func printHistogram(histogram map[int]int, maxVal int) {
	fmt.Println(strings.Repeat("=", maxVal))
	for i := 2; i <= maxVal; i++ {
		s := strings.Repeat("*", histogram[i])
		fmt.Printf("%3v: %v %v\n", i, s, histogram[i])
	}
}

func createHistogramRecursion(dice []int) map[int]int {
	if len(dice) == 0 {
		return map[int]int{0: 1}
	}

	histogram := map[int]int{}
	for i := 1; i <= dice[0]; i++ {
		results := createHistogramRecursion(dice[1:])
		for sum, occurrences := range results {
			histogram[sum+i] += occurrences
		}
	}

	return histogram
}

func createHistogram(dice []int) map[int]int {
	histogram := map[int]int{0: 1}

	for _, sides := range dice {
		results := map[int]int{}
		for i := 1; i <= sides; i++ {
			for sum, occurrences := range histogram {
				results[sum+i] += occurrences
			}
		}
		histogram = results
	}

	return histogram
}
