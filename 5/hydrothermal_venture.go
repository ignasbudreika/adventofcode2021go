package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type vent struct {
	StartX     int
	StartY     int
	EndX       int
	EndY       int
	Horizontal bool
	Vertical   bool
	Diagonal   bool
}

func main() {
	part1()
	part2()
}

func part1() {
	lines, err := parse("input.txt")
	if err != nil {
		log.Fatalf("reading input went :( (%v)", err)
	}

	diagram := make(map[int]map[int]int, 0)
	collisions := 0

	for _, line := range lines {
		if line.Horizontal {
			collisions = handleHorizontal(line, diagram, collisions)
		} else if line.Vertical {
			collisions = handleVertical(line, diagram, collisions)
		}
	}

	fmt.Println(collisions)
}

func part2() {
	lines, err := parse("input.txt")
	if err != nil {
		log.Fatalf("reading input went :( (%v)", err)
	}

	diagram := make(map[int]map[int]int, 0)
	collisions := 0

	for _, line := range lines {
		if line.Horizontal {
			collisions = handleHorizontal(line, diagram, collisions)
		} else if line.Vertical {
			collisions = handleVertical(line, diagram, collisions)
		} else {
			collisions = handleDiagonal(line, diagram, collisions)
		}
	}

	fmt.Println(collisions)
}

func handleHorizontal(line vent, diagram map[int]map[int]int, collisions int) int {
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

	return collisions
}

func handleVertical(line vent, diagram map[int]map[int]int, collisions int) int {
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

	return collisions
}

func handleDiagonal(line vent, diagram map[int]map[int]int, collisions int) int {
	xDirection := 1
	yDirection := 1
	j := line.StartY

	if line.StartX > line.EndX {
		xDirection = -1
	}

	if line.StartY > line.EndY {
		yDirection = -1
	}

	for i := line.StartX; i-xDirection != line.EndX; i += xDirection {
		if _, ok := diagram[i]; ok {
			if _, ok := diagram[i][j]; ok {
				if diagram[i][j] == 1 {
					diagram[i][j] = 2
					collisions++
				}
			} else {
				diagram[i][j] = 1
			}
		} else {
			diagram[i] = make(map[int]int)
			diagram[i][j] = 1
		}

		j += yDirection
	}

	return collisions
}

func parse(filename string) ([]vent, error) {
	input, err := os.ReadFile(filename)
	if err != nil {
		return []vent{}, err
	}

	lines := make([]vent, 0)
	for _, l := range strings.Split(string(input), "\n") {
		coords := strings.Split(l, " -> ")
		start := strings.Split(coords[0], ",")
		end := strings.Split(coords[1], ",")
		startX, _ := strconv.Atoi(start[0])
		startY, _ := strconv.Atoi(start[1])
		endX, _ := strconv.Atoi(end[0])
		endY, _ := strconv.Atoi(end[1])

		horizontal := false
		vertical := false
		diagonal := false
		if startX != endX && startY != endY {
			horizontal = false
			vertical = false
			diagonal = true
		} else if startX != endX {
			horizontal = true
			vertical = false
			diagonal = false
		} else {
			horizontal = false
			vertical = true
			diagonal = false
		}

		lines = append(lines, vent{StartX: startX, StartY: startY, EndX: endX, EndY: endY,
			Horizontal: horizontal, Vertical: vertical, Diagonal: diagonal})
	}

	return lines, nil
}
