package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	/* generateTestCases("test_data.txt")
	testcases, err := readTestData("test_data.txt")
	if err == nil {
		for _, problem := range testcases {
			if NaiveDPKnapsack(problem.Weights, problem.Capacity) != problem.Answer {
				fmt.Println(problem)
			}
			//fmt.Println(BruteForceKnapsack(problem.Weights, problem.Capacity))
		}
	}*/

	// Calculate the result
	//fmt.Println(NaiveDPKnapsack(weight, W))
	//fmt.Println(bruteForceKnapsack(weight, W))
	tests, err := readTestData("test_data.txt")
	if err != nil {
		fmt.Println("Error reading test data")
	}
	for _, test := range tests {
		// time how long the SpaceEfficientKnapsack function takes
		start := time.Now()
		SpaceEfficientKnapsack(test.Weights, test.Capacity)
		time1 := time.Since(start)
		//print
		start = time.Now()
		ans := ApproximateSolution(test.Weights, test.Capacity, 0.1, 0.1)
		time2 := time.Since(start)
		//float version of the next line
		err := 1 - math.Abs(float64(test.Answer)-float64(ans))/float64(test.Answer)
		if err == 2 {
			fmt.Println("error:", err)
		}

		//difference in time as a percentage of time 1
		fmt.Println("time1: ", time1, "time2: ", time2)

	}

	var memStats runtime.MemStats

	// Read current memory statistics
	runtime.ReadMemStats(&memStats)

	// Print heap memory statistics
	// fmt.Printf("Heap Alloc: %v bytes\n", memStats.HeapAlloc)
	// fmt.Printf("Heap Sys: %v bytes\n", memStats.HeapSys)
	// fmt.Printf("Heap Idle: %v bytes\n", memStats.HeapIdle)
	// fmt.Printf("Heap Inuse: %v bytes\n", memStats.HeapInuse)
	// fmt.Printf("Heap Released: %v bytes\n", memStats.HeapReleased)
	// fmt.Printf("Heap Objects: %v\n", memStats.HeapObjects)
}
