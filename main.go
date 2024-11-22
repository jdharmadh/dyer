package main

import (
	"fmt"
	"runtime"
)

func main() {
	generateTestCases("test_data.txt") /*
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

	var memStats runtime.MemStats

	// Read current memory statistics
	runtime.ReadMemStats(&memStats)

	// Print heap memory statistics
	fmt.Printf("Heap Alloc: %v bytes\n", memStats.HeapAlloc)
	fmt.Printf("Heap Sys: %v bytes\n", memStats.HeapSys)
	fmt.Printf("Heap Idle: %v bytes\n", memStats.HeapIdle)
	fmt.Printf("Heap Inuse: %v bytes\n", memStats.HeapInuse)
	fmt.Printf("Heap Released: %v bytes\n", memStats.HeapReleased)
	fmt.Printf("Heap Objects: %v\n", memStats.HeapObjects)
}
