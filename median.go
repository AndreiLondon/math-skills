package main

import (
	"sort"
)

// https://gosamples.dev/calculate-median/#:~:text=To%20calculate%20the%20median%20in,value%20from%20this%20sorted%20slice.

func getMedian(result []float64) float64 {
	myMap := make([]float64, len(result))
	copy(myMap, result)

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

	return median
}
