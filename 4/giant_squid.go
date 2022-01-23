package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingoBoard struct {
	Rows [][]int
	Sum  int
}

type result struct {
	Won bool
	Sum int
}

type bingoGame struct {
	Numbers []int
	Boards  []bingoBoard
}

func main() {
	part1()
}

func part1() {
	game, err := parse("input.txt")
	if err != nil {
		log.Fatalf("bingo subsystem has failed: %v", err)
	}

	for _, number := range game.Numbers {
		marked := make(chan result, len(game.Boards))

		for i := 0; i < len(game.Boards); i++ {
			go markNumber(marked, &game.Boards[i], number)
		}

		stop := false

		for i := 0; i < len(game.Boards); i++ {
			res := <-marked
			if res.Won {
				stop = true
				fmt.Println(number * res.Sum)
			}
		}

		if stop {
			break
		}
	}
}

func markNumber(marked chan<- result, board *bingoBoard, number int) {
	found := false

	// mark number
	for i := 0; i < len(board.Rows); i++ {
		for j := 0; j < len(board.Rows[i]); j++ {
			if board.Rows[i][j] == number {
				board.Rows[i][j] = 0
				board.Sum -= number
				found = true
				goto checkIfBoardWon
			}
		}
	}

checkIfBoardWon:
	if !found {
		marked <- result{Won: false, Sum: board.Sum}
		return
	}

	// calc row sum
	for _, row := range board.Rows {
		sum := 0

		for _, n := range row {
			sum += n
		}

		if sum == 0 {
			marked <- result{Won: true, Sum: board.Sum}
			return
		}
	}

	// calc column sum
	for i := 0; i < len(board.Rows[0]); i++ {
		sum := 0

		for _, row := range board.Rows {
			sum += row[i]
		}

		if sum == 0 {
			marked <- result{Won: true, Sum: board.Sum}
			return
		}
	}

	marked <- result{Won: false, Sum: board.Sum}
}

func parse(filename string) (bingoGame, error) {
	numbers := make([]int, 0)
	boards := make([]bingoBoard, 0)

	input, err := os.ReadFile(filename)
	if err != nil {
		return bingoGame{}, fmt.Errorf("uh oh reading bingo game data went wrong: %v", err)
	}

	lines := strings.Split(string(input), "\n")

	if len(lines) <= 1 {
		return bingoGame{}, errors.New("huh where's the bingo game")
	}

	for _, number := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(number)

		numbers = append(numbers, n)
	}

	for i := 2; i < len(lines); i += 6 {
		rows := make([][]int, 0)
		sum := 0

		for j := i; j < i+5; j++ {
			row := make([]int, 0)

			for _, number := range strings.Fields(lines[j]) {
				n, _ := strconv.Atoi(number)

				row = append(row, n)
				sum += n
			}

			rows = append(rows, row)
		}

		boards = append(boards, bingoBoard{Rows: rows, Sum: sum})
	}

	return bingoGame{Boards: boards, Numbers: numbers}, nil
}
