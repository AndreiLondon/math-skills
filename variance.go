package main

import "math"

func getVariance(data []float64) float64 {
	sumSq := 0.0
	meanFloat64 := getAverage(data)
	for _, num := range data {
		numFloat64 := float64(num)
		sumSq += (numFloat64 - meanFloat64) * (numFloat64 - meanFloat64)
	}
	lenFloat64 := float64(len(data))
	result := float64(sumSq) / lenFloat64

	return math.Round(result)

}
