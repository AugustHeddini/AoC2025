package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func print_list(list []int) {
	for i := range list {
		fmt.Print(list[i], ", ")
	}
	fmt.Println()
}

func count_past_zero(moves []int, curr int) int {
	count := 0
	for _, move := range moves {
		curr += move
		curr %= 100
		if curr == 0 {
			count++
		}
	}
	return count
}

func count_all_zeroes(moves []int, curr int) int {
	count := 0
	for _, move := range moves {
		if abs(move) >= 100 {
			count += abs(move) / 100
			move -= (move / 100) * 100
		}
		if (curr < 0 && curr + move >= 0) ||
			(curr > 0 && curr + move <= 0) {
			count++
		}
		if abs(curr + move) >= 100 {
			count++
		}

		curr += move
		curr %= 100
	}
	return count
}

func parse_input(input *os.File) []int {
	line_scanner := bufio.NewScanner(input)

	moves := []int{}
	for line_scanner.Scan() {
		line := line_scanner.Text()
		if line == "" {
			continue
		}

		neg := false
		if line[0] == 'L' {
			neg = true
		}

		if i, err := strconv.Atoi(line[1:]); err == nil {
			if neg {
				i = -i
			}
			moves = append(moves, i)
		} else {
			panic(err)
		}
	}

	return moves
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	moves := parse_input(input)

	// res := count_past_zero(moves, 50)
	res := count_all_zeroes(moves, 50)

	fmt.Println("Needle passed zero ", res, " times")
}
