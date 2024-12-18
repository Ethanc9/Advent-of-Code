package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readInput("Day-2/input.txt")
	if err != nil {
		panic(err)
	}

	count := countSafeNumbers(lines)
	println("Number of safe sequences:", count)
}

func countSafeNumbers(lines []string) int {
	safeCount := 0

	for _, line := range lines {
		numbers := strings.Fields(line)
		prevNum := -1
		isSafe := true
		isIncreasing := true  // Track direction
		directionSet := false 

		for i, numStr := range numbers {
			num, _ := strconv.Atoi(numStr)

			if i > 0 {
				diff := num - prevNum
				// Set direction based on first pair
				if !directionSet {
					isIncreasing = diff > 0
					directionSet = true
				} else {
					// Check if direction remains consistent
					if (diff > 0) != isIncreasing {
						isSafe = false
					}
				}
				
				// Check if absolute difference is between 1-3
				if abs(diff) < 1 || abs(diff) > 3 {
					isSafe = false
				}
			}
			prevNum = num
		}

		if isSafe {
			safeCount++
		}
	}

	return safeCount
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
