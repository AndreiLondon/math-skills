package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*The purpose of this project is for you to calculate the following:

Average
Median
Variance
Standard Deviation

Your program must be able to read from a file and print the result of each statistic mentioned above.
In other words, your program must be able to read the data present in the path passed as argument.
*/

func getArguments() string {
	// ignore item 0 and get everytning that follows item 0
	args := os.Args[1:]

	if len(args) == 0 || len(args) > 1 {
		fmt.Println("invalid number of arguments")
		os.Exit(0)
	}
	// if len(args) == 1 {
	// 	return args[0]
	// }

	return args[0]
}

func readFile(file string) []float64 {
	data, err := os.ReadFile(file)
	dataArr := strings.Split(string(data), "\n")
	result := make([]float64, 0)
	for i := 0; i < len(dataArr); i++ {
		conv, _ := strconv.Atoi(dataArr[i])
		result = append(result, float64(conv))
	}

	//result := binary.BigEndian.Uint64(dataArr[i])
	if err != nil {
		log.Fatal("Cannot read the file")
		os.Exit(0)
	}
	return result
}
