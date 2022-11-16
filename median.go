package main

import (
	"math"
	"sort"
)

// https://gosamples.dev/calculate-median/#:~:text=To%20calculate%20the%20median%20in,value%20from%20this%20sorted%20slice.

func getMedian(data []float64) float64 {
	myMap := make([]float64, len(data))
	copy(myMap, data)
	sort.Float64s(myMap)
	var median float64
	l := len(myMap)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (myMap[l/2-1] + myMap[l/2]) / 2
	} else {
		median = myMap[l/2]
	}

	return math.Round(median)
}
