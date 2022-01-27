package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type line struct {
	StartX     int
	StartY     int
	EndX       int
	EndY       int
	Horizontal bool
}

func main() {
	lines, err := parse("input.txt")
	if err != nil {
		log.Fatalf("reading input went :( (%v)", err)
	}

	diagram := make(map[int]map[int]int, 0)
	collisions := 0

	for _, line := range lines {
		if line.Horizontal {
			from := line.StartX
			to := line.EndX
			if line.StartX > line.EndX {
				from = line.EndX
				to = line.StartX
			}

			for i := from; i <= to; i++ {
				if column, ok := diagram[i]; ok {
					if val, ok := column[line.StartY]; ok {
						if val == 1 {
							diagram[i][line.StartY] = 2
							collisions++
						}
					} else {
						diagram[i][line.StartY] = 1
					}
				} else {
					diagram[i] = make(map[int]int)
					diagram[i][line.StartY] = 1
				}
			}
		} else {
			from := line.StartY
			to := line.EndY
			if line.StartY > line.EndY {
				from = line.EndY
				to = line.StartY
			}

			column, ok := diagram[line.StartX]
			if !ok {
				diagram[line.StartX] = make(map[int]int)
			}

			for i := from; i <= to; i++ {
				if val, ok := column[i]; ok {
					if val == 1 {
						diagram[line.StartX][i] = 2
						collisions++
					}
				} else {
					diagram[line.StartX][i] = 1
				}
			}
		}
	}

	fmt.Println(collisions)
}

func parse(filename string) ([]line, error) {
	input, err := os.ReadFile(filename)
	if err != nil {
		return []line{}, err
	}

	lines := make([]line, 0)
	for _, l := range strings.Split(string(input), "\n") {
		coords := strings.Split(l, " -> ")
		start := strings.Split(coords[0], ",")
		end := strings.Split(coords[1], ",")
		startX, _ := strconv.Atoi(start[0])
		startY, _ := strconv.Atoi(start[1])
		endX, _ := strconv.Atoi(end[0])
		endY, _ := strconv.Atoi(end[1])

		if startX != endX && startY != endY {
			continue
		}

		horizontal := false
		if startX != endX {
			horizontal = true
		}

		lines = append(lines, line{StartX: startX, StartY: startY,
			EndX: endX, EndY: endY, Horizontal: horizontal})
	}

	return lines, nil
}
