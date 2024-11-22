package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type testCase struct {
	Capacity uint
	Weights  []uint
	Answer   uint
}

// Function to read test data and answers from a .txt file
func readTestData(filename string) ([]testCase, error) {
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
	var testCases []testCase

	// Read each test case
	for i := 0; i < numTestCases; i++ {
		// Read the first line (length of weights, capacity, and answer)
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		numWeights, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		capacity, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		answer, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, err
		}

		// Read the weights line
		scanner.Scan()
		weightsLine := scanner.Text()
		weightsParts := strings.Fields(weightsLine)
		var weights []uint
		for _, weightStr := range weightsParts {
			weight, err := strconv.Atoi(weightStr)
			if err != nil {
				return nil, err
			}
			weights = append(weights, uint(weight))
		}

		// Add the test case to the slice
		if len(weights) != numWeights {
			return nil, fmt.Errorf("mismatched weights count: expected %d, got %d", numWeights, len(weights))
		}
		testCases = append(testCases, testCase{Capacity: uint(capacity), Weights: weights, Answer: uint(answer)})
	}

	return testCases, nil
}

// Function to write test data and answers to a .txt file
func writeTestData(filename string, testCases []testCase) error {
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
		// Write the length of the weights table, capacity, and the answer
		_, err = file.WriteString(fmt.Sprintf("%d %d %d\n", len(testCase.Weights), testCase.Capacity, testCase.Answer))
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
