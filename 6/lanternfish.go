package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type lanternfishes struct {
	Fishes map[int]int
	Total  int
}

func main() {
	lanternfishes, err := parse("input.txt")
	if err != nil {
		log.Fatalf("catching fishes failed: %v", err)
	}

	for day := 1; day <= 80; day++ {
		ready := lanternfishes.Fishes[0]

		for i := 0; i < 8; i++ {
			lanternfishes.Fishes[i] = lanternfishes.Fishes[i+1]
		}

		lanternfishes.Fishes[6] += ready
		lanternfishes.Fishes[8] = ready

		lanternfishes.Total += ready
	}

	fmt.Println(lanternfishes.Total)
}

func parse(filename string) (lanternfishes, error) {
	input, err := os.ReadFile(filename)
	if err != nil {
		return lanternfishes{}, err
	}

	days := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	total := 0
	for _, number := range strings.Split(string(input), ",") {
		n, _ := strconv.Atoi(number)

		days[n]++
		total++
	}

	return lanternfishes{Fishes: days, Total: total}, nil
}
