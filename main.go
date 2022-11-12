package main

import "fmt"

func main() {

	argument := getArguments()

	file := readFile(argument)

	fmt.Printf("Average: %v", getAverage)
	fmt.Printf("Median: %v", getMedian)
	fmt.Printf("Variance: %v", getVariance)
	fmt.Printf("Standard Deviation: %v", getStDeviation)

}
