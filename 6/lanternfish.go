package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type lanternfishes struct {
	Fishes map[int]uint64
	Total  uint64
}

func main() {
	data, err := parse("input.txt")
	if err != nil {
		log.Fatalf("catching fishes failed: %v", err)
	}

	// something does not work when copying the map:///
	// dataCopy := lanternfishes{
	// 	Fishes: map[int]uint64{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0},
	// 	Total:  0,
	// }
	// for day, fishes := range data.Fishes {
	// 	fmt.Println(day, fishes)
	// 	dataCopy.Fishes[day] = fishes
	// }

	// fmt.Printf("after 80 days: %d\n", calcFishes(data, 80))
	fmt.Printf("after 256 days: %d\n", calcFishes(data, 256))
}

func calcFishes(data lanternfishes, days int) uint64 {
	for day := 1; day <= days; day++ {
		ready := data.Fishes[0]

		for i := 0; i < 8; i++ {
			data.Fishes[i] = data.Fishes[i+1]
		}

		data.Fishes[6] += ready
		data.Fishes[8] = ready

		data.Total += ready
	}

	return data.Total
}

func parse(filename string) (lanternfishes, error) {
	input, err := os.ReadFile(filename)
	if err != nil {
		return lanternfishes{}, err
	}

	days := map[int]uint64{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	total := 0
	for _, number := range strings.Split(string(input), ",") {
		n, _ := strconv.Atoi(number)

		days[n]++
		total++
	}

	return lanternfishes{Fishes: days, Total: uint64(total)}, nil
}
