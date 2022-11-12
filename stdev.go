package main

import "math"

func getStDeviation(result []float64, mean float64) float64 {

	//func stdDev(numbers []float64, mean float64) float64 {
	total := 0.0
	for _, number := range result {
		total += math.Pow(number-mean, 2)
	}
	// variance := total / float64(len(result)-1)
	variance := total / float64(len(result))
	return math.Sqrt(variance)
}
