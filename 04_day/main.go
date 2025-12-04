package main

import (
	"bufio"
	"fmt"
	"os"
)

var dirs = [8]pos{ 
	{ -1, -1 },  { -1, 0 },  { -1, 1 },
	{ 0, -1 },/* { 0, 0 } */ { 0, 1 },
	{ 1, -1 }, { 1, 0 }, { 1, 1 } }

type pos struct {
	x int
	y int
}

func (self pos) add(other pos) pos {
	return pos{ self.x + other.x, self.y + other.y }
}

func add_neighbours(nb *map[pos]int, center pos) {
	for _, dir := range dirs {
		(*nb)[center.add(dir)] += 1
	}
}

func remove_neighbours(nb *map[pos]int, center pos) {
	for _, dir := range dirs {
		(*nb)[center.add(dir)] -= 1
	}
}

func parse_paper_rolls(input *os.File) ([]pos, map[pos]int) {
	
	scanner := bufio.NewReader(input)
	neighbours := make(map[pos]int)
	paper_rolls := []pos{}

	x := 0
	y := 0
	for {
		byte, err := scanner.ReadByte()
		if err != nil {
			break
		}
		x += 1
		if byte == '.' {
			continue
		}
		if byte == '\n' {
			y += 1
			x = 0
			continue
		}

		curr_pos := pos { x, y }
		paper_rolls = append(paper_rolls, curr_pos)
		add_neighbours(&neighbours, curr_pos)
	}

	return paper_rolls, neighbours
}

func count_accessible(papers *[]pos, nb *map[pos]int) (count int) {

	for _, paper := range *papers {
		if (*nb)[paper] < 4 {
			count++
		}
	}
	return
}

func count_all_accessible(papers *[]pos, nb map[pos]int) (count int) {

	var res int
	to_check := *papers
	for {
		res = 0
		remaining := []pos{}
		for _, paper := range to_check {
			if nb[paper] < 4 {
				res++
				remove_neighbours(&nb, paper)
			} else {
				remaining = append(remaining, paper)
			}
		}
		to_check = remaining

		if res == 0 {
			break
		}
		count += res
	}
	return
}


func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	papers, nb := parse_paper_rolls(input)

	// res := count_accessible(&papers, &nb)
	res := count_all_accessible(&papers, nb)

	fmt.Println("Found ", res, " accessible paper rolls")
}
