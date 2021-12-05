package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	measurements, err := parseInput("input.txt")
	if err != nil {
		log.Fatalf("input error: %v", err)
	}

	increased := 0
	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			increased++
		}
	}

	fmt.Println(increased)
}

func part2() {
	measurements, err := parseInput("input.txt")
	if err != nil {
		log.Fatalf("input error: %v", err)
	}

	increased := 0
	for i := 3; i < len(measurements); i++ {
		if measurements[i] > measurements[i-3] {
			increased++
		}
	}

	fmt.Println(increased)
}

func parseInput(filename string) ([]int, error) {
	measurements := make([]int, 0)

	input, err := os.ReadFile(filename)
	if err != nil {
		return measurements, fmt.Errorf("reading input went wrong: %v", err)
	}

	for _, line := range strings.Split(string(input), "\n") {
		if len(line) > 0 {
			depth, err := strconv.ParseInt(line, 10, 32)
			if err != nil {
				return measurements, fmt.Errorf("that was not a number: %v", err)
			}

			measurements = append(measurements, int(depth))
		}
	}

	return measurements, nil
}
