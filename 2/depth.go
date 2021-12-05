package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	forward = iota
	down
	up
)

type Command struct {
	Direction int
	Units     int32
}

type Location struct {
	Horizontal int
	Depth      int
	Aim        int
}

func main() {
	part1()
	part2()
}

func part1() {
	commands, err := parseInput("input.txt")
	if err != nil {
		log.Fatalf("input error: %v", err)
	}

	submarine := Location{Horizontal: 0, Depth: 0}

	for _, c := range commands {
		if c.Direction == forward {
			submarine.Horizontal += int(c.Units)
		} else if c.Direction == down {
			submarine.Depth += int(c.Units)
		} else {
			submarine.Depth -= int(c.Units)
		}
	}

	fmt.Println(submarine.Horizontal * submarine.Depth)
}

func part2() {
	commands, err := parseInput("input.txt")
	if err != nil {
		log.Fatalf("input error: %v", err)
	}

	submarine := Location{Horizontal: 0, Depth: 0}

	for _, c := range commands {
		if c.Direction == forward {
			submarine.Horizontal += int(c.Units)
			submarine.Depth += submarine.Aim * int(c.Units)
		} else if c.Direction == down {
			submarine.Aim += int(c.Units)
		} else {
			submarine.Aim -= int(c.Units)
		}
	}

	fmt.Println(submarine.Horizontal * submarine.Depth)
}

func parseInput(filename string) ([]Command, error) {
	commands := make([]Command, 0)

	input, err := os.ReadFile(filename)
	if err != nil {
		return commands, fmt.Errorf("reading input went wrong: %v", err)
	}

	for _, line := range strings.Split(string(input), "\n") {
		if len(line) > 0 {
			var direction int
			switch strings.Split(line, " ")[0] {
			case "forward":
				direction = forward
			case "down":
				direction = down
			case "up":
				direction = up
			default:
				return commands, fmt.Errorf("unclear command: %s", strings.Split(line, " ")[0])
			}

			units, err := strconv.ParseInt(strings.Split(line, " ")[1], 10, 64)
			if err != nil {
				return commands, fmt.Errorf("that was not a number: %v", err)
			}

			commands = append(commands, Command{Direction: direction, Units: int32(units)})
		}
	}

	return commands, nil
}
