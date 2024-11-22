package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to read test data from a .txt file
func readTestData(filename string) ([]struct {
	Capacity int
	Weights  []int
}, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the number of test cases
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	numTestCases, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}

	// Create a slice to store the test cases
	var testCases []struct {
		Capacity int
		Weights  []int
	}

	// Read each test case
	for i := 0; i < numTestCases; i++ {
		// Read the first line (length of weights and capacity)
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		/*numWeights, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}*/
		capacity, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		// Read the weights line
		scanner.Scan()
		weightsLine := scanner.Text()
		weightsParts := strings.Fields(weightsLine)
		var weights []int
		for _, weightStr := range weightsParts {
			weight, err := strconv.Atoi(weightStr)
			if err != nil {
				return nil, err
			}
			weights = append(weights, weight)
		}

		// Add the test case to the slice
		testCases = append(testCases, struct {
			Capacity int
			Weights  []int
		}{Capacity: capacity, Weights: weights})
	}

	return testCases, nil
}

func read() {
	// Read test data from "test_data.txt"
	testCases, err := readTestData("test_data.txt")
	if err != nil {
		fmt.Println("Error reading data:", err)
	} else {
		fmt.Println("Test data read successfully.")
		for _, testCase := range testCases {
			fmt.Printf("Capacity: %d, Weights: %v\n", testCase.Capacity, testCase.Weights)
		}
	}
}

// Function to write test data to a .txt file
func writeTestData(filename string, testCases []struct {
	Capacity int
	Weights  []int
}) error {
	// Create or open the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the number of test cases
	_, err = file.WriteString(fmt.Sprintf("%d\n", len(testCases)))
	if err != nil {
		return err
	}

	// Write data for each test case
	for _, testCase := range testCases {
		// Write the length of the weights table and the capacity
		_, err = file.WriteString(fmt.Sprintf("%d %d\n", len(testCase.Weights), testCase.Capacity))
		if err != nil {
			return err
		}

		// Write the weights
		_, err = file.WriteString(fmt.Sprintf("%s\n", strings.Trim(fmt.Sprint(testCase.Weights), "[]")))
		if err != nil {
			return err
		}
	}

	return nil
}

func write() {
	// Example test data
	testCases := []struct {
		Capacity int
		Weights  []int
	}{
		{40, []int{10, 20, 30, 40}},
		{50, []int{5, 10, 15, 20, 25}},
	}

	// Write test data to "test_data.txt"
	err := writeTestData("test_data.txt", testCases)
	if err != nil {
		fmt.Println("Error writing data:", err)
	} else {
		fmt.Println("Test data written successfully.")
	}
}
