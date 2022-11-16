package main

import "math"

func getAverage(data []float64) float64 {
	array := 0.0
	for _, v := range data {
		array += v
	}
	result := array / float64(len(data))
	return math.Round(result)

}
