package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type message struct {
	Index           int
	MostCommonByte  byte
	LeastCommonByte byte
}

func main() {
	part1()
	part2()
}

func part1() {
	numbers, err := parse("input.txt")
	if err != nil {
		log.Fatal(":(")
	}

	if len(numbers) == 0 {
		fmt.Println("0")
		os.Exit(1)
	}

	answers := make(chan message, len(numbers[0]))
	wg := &sync.WaitGroup{}

	for i := 0; i < len(numbers[0]); i++ {
		wg.Add(1)
		go countOnes(wg, numbers, i, answers)
	}

	wg.Wait()

	gammaBytes := make([]string, len(numbers[0]))
	epsilonBytes := make([]string, len(numbers[0]))

	for i := 0; i < len(numbers[0]); i++ {
		msg := <-answers
		gammaBytes[msg.Index] = string(msg.MostCommonByte)
		epsilonBytes[msg.Index] = string(msg.LeastCommonByte)
	}

	gammaRate, err := strconv.ParseInt(strings.Join(gammaBytes, ""), 2, 64)
	if err != nil {
		log.Fatalf("something wrong with gamma rate: %s", strings.Join(gammaBytes, ""))
	}

	epsilonRate, err := strconv.ParseInt(strings.Join(epsilonBytes, ""), 2, 64)
	if err != nil {
		log.Fatalf("something wrong with epsilon rate: %s", strings.Join(epsilonBytes, ""))
	}

	fmt.Printf("gamma rate: %d\n", gammaRate)
	fmt.Printf("epsilon rate: %d\n", epsilonRate)
	fmt.Printf("power consumption: %d\n", gammaRate*epsilonRate)
}

func part2() {
	numbers, err := parse("input.txt")
	if err != nil {
		log.Fatal(":(")
	}

	if len(numbers) == 0 {
		fmt.Println("0")
		os.Exit(1)
	}

	oxygenGeneratorBytes := findOxygenGeneratorRatingRecursively(numbers, 0)
	CO2ScrubberBytes := findCO2SCrubberRatingRecursively(numbers, 0)

	oxygenGeneratorRating, err := strconv.ParseInt(oxygenGeneratorBytes, 2, 64)
	if err != nil {
		log.Fatalf("something wrong with oxygen generator rating: %s", oxygenGeneratorBytes)
	}

	CO2ScrubberRating, err := strconv.ParseInt(CO2ScrubberBytes, 2, 64)
	if err != nil {
		log.Fatalf("something wrong with co2 scrubber rating: %s", CO2ScrubberBytes)
	}

	fmt.Printf("oxygen generator rating: %d\n", oxygenGeneratorRating)
	fmt.Printf("co2 scrubber rating: %d\n", CO2ScrubberRating)
	fmt.Printf("power consumption: %d\n", oxygenGeneratorRating*CO2ScrubberRating)
}

func countOnes(wg *sync.WaitGroup, numbers []string, i int, answer chan<- message) {
	defer wg.Done()
	count := 0

	for _, number := range numbers {
		n, _ := strconv.Atoi(string(number[i]))

		count += n
	}

	if count > len(numbers)/2 {
		answer <- message{Index: i, MostCommonByte: '1', LeastCommonByte: '0'}
	} else {
		answer <- message{Index: i, MostCommonByte: '0', LeastCommonByte: '1'}
	}
}

func findOxygenGeneratorRatingRecursively(numbers []string, i int) string {
	filteredWithOnes := make([]string, 0)
	filteredWithZeros := make([]string, 0)

	for _, number := range numbers {
		if number[i] == '1' {
			filteredWithOnes = append(filteredWithOnes, number)
		} else {
			filteredWithZeros = append(filteredWithZeros, number)
		}
	}

	if len(filteredWithOnes) >= len(filteredWithZeros) {
		if len(filteredWithOnes) == 1 {
			return filteredWithOnes[0]
		}

		return findOxygenGeneratorRatingRecursively(filteredWithOnes, i+1)
	}

	if len(filteredWithZeros) == 1 {
		return filteredWithZeros[0]
	}
	return findOxygenGeneratorRatingRecursively(filteredWithZeros, i+1)
}

func findCO2SCrubberRatingRecursively(numbers []string, i int) string {
	filteredWithOnes := make([]string, 0)
	filteredWithZeros := make([]string, 0)

	for _, number := range numbers {
		if number[i] == '1' {
			filteredWithOnes = append(filteredWithOnes, number)
		} else {
			filteredWithZeros = append(filteredWithZeros, number)
		}
	}

	if len(filteredWithOnes) >= len(filteredWithZeros) {
		if len(filteredWithZeros) == 1 {
			return filteredWithZeros[0]
		}

		return findCO2SCrubberRatingRecursively(filteredWithZeros, i+1)
	}

	if len(filteredWithOnes) == 1 {
		return filteredWithOnes[0]
	}
	return findCO2SCrubberRatingRecursively(filteredWithOnes, i+1)
}

func parse(filename string) ([]string, error) {
	numbers := make([]string, 0)

	input, err := os.ReadFile(filename)
	if err != nil {
		return numbers, fmt.Errorf("uh oh reading diagnose went wrong: %v", err)
	}

	for _, line := range strings.Split(string(input), "\n") {
		numbers = append(numbers, line)
	}

	return numbers, nil
}
