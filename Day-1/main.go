package main

import (
    "fmt"
    "sort"
    "strings"
    "strconv"
    "os"
    "bufio"
)

func main() {
    // Read input from file
    file, err := os.Open("Day-1/input.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()

    leftList := make([]int, 0)
    rightList := make([]int, 0)

    // Parse input file
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            continue
        }
        
        numbers := strings.Fields(line)
        if len(numbers) != 2 {
            continue
        }
        
        left, _ := strconv.Atoi(numbers[0])
        right, _ := strconv.Atoi(numbers[1])
        
        leftList = append(leftList, left)
        rightList = append(rightList, right)
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }

    // Sort both lists
    sort.Ints(leftList)
    sort.Ints(rightList)

    // Calculate total distance
    totalDistance := 0
    for i := 0; i < len(leftList); i++ {
        distance := abs(leftList[i] - rightList[i])
        totalDistance += distance
    }

    fmt.Printf("Total distance: %d\n", totalDistance)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}