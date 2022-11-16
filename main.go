package main

import (
	"fmt"
	"math"
)

func main() {

	argument := getArguments()

	file := readFile(argument)

	average := getAverage(file)
	median := getMedian(file)
	variance := getVariance(file)
	stDeviation := getStDeviation(file, average)

	fmt.Printf("Average: %v \n", math.Round(average))
	fmt.Printf("Median: %v \n", math.Round(median))
	fmt.Printf("Variance: %v \n", math.Round(variance))
	fmt.Printf("Standard Deviation: %v \n", math.Round(stDeviation))

}
