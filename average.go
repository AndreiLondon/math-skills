package main

func getAverage(result []float64) float64 {
	array := 0.0
	for _, v := range result {
		array += v
	}
	return array / float64(len(result))

}
