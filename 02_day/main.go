package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type interval struct {
	start int
	end int
}

func digits(i int) (digits int) {
	if i == 0 {
		return 1
	}
	digits = 0
	for i != 0 {
		i /= 10
		digits++
	}
	return
}

func reduce(val int, digits int) int {
	for range digits {
		val /= 10
	}
	return val
}

func pow(base, pow int) (res int) {
	res = 1
	for range pow {
		res *= base
	}
	return
}

func parse_intervals(input *os.File) (intervals []interval) {
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	line := scanner.Text()
	line = strings.Trim(line, " \n")

	interval_defs := strings.Split(line, ",")

	for _, interval_string := range interval_defs {
		ends := strings.Split(interval_string, "-")
		if start, err := strconv.Atoi(ends[0]); err == nil {
			if end, err := strconv.Atoi(ends[1]); err == nil {
				intervals = append(intervals, interval{start, end})
			}
		} else {
			panic(err)
		}
	}
	return
}

func count_all_invalids(intervals []interval) (sum int) {
	sum = 0

	for _, iv := range intervals {
		for i := iv.start; i <= iv.end; i++ {
			digits := digits(i)

			segments:
			for seg := 1; seg <= digits/2; seg++ {

				part := reduce(i, digits - seg)
				step := pow(10, seg)
				num := part
				for num < i {
					num = num * step + part
				}

				if num == i {
					sum += i
					break segments
				}
			}
		}
	}

	return
}

func count_invalids(intervals []interval) (sum int) {
	sum = 0

	for _, iv := range intervals {
		for i := iv.start; i <= iv.end; i++ {
			digits := digits(i)
			if digits % 2 != 0 {
				continue
			}
			digits /= 2

			half := reduce(i, digits)

			if half * pow(10, digits) + half == i {
				sum += i
			}
		}
	}

	return
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	intervals := parse_intervals(input)
	fmt.Println(intervals)

	// res := count_invalids(intervals)
	res := count_all_invalids(intervals)

	fmt.Println("Result: ", res)
}
	

