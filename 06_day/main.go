package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Op func([]int) int

type problem struct {
	values []int
	operator Op
}

func (p *problem) operate() int {
	return p.operator(p.values)
}

func mul(l []int) int {
	prod := 1
	for _, v := range l {
		prod *= v
	}
	return prod
}

func sum(l []int) (total int) {
	for _, v := range l {
		total += v
	}
	return
}

func filter_whitespace(strs []string) []string {
	keep := make([]string, 0)
	for _, s := range strs {
		if !(s == "" || s == " ") {
			keep = append(keep, s)
		}
	}
	return keep
}

func isolate_digits(i int) []int {
	digits := []int{}
	for i > 0 {
		digits = append(digits, i % 10)
		i /= 10
	}
	return digits
}

func parse_problems(input *os.File) []problem {

	scanner := bufio.NewScanner(input)
	problems := make([]problem, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		values := strings.Split(line, " ")
		values = filter_whitespace(values)
		for i, v := range values {
			if v_int, err := strconv.Atoi(v); err == nil {
				if i >= len(problems) {
					problems = append(problems, problem{ []int{v_int}, mul })
					continue
				}
				problems[i].values = append(problems[i].values, v_int)
			} else {
				if v == "+" {
					problems[i].operator = sum
				}
			}
		}
	}
	return problems
}

func parse_problems_vertical(input *os.File) []problem {

	reader := bufio.NewReader(input)
	problems := make([]problem, 0)

	i := 0
	dig_i := 0
	spaces := 1
	for {
		char, err := reader.ReadByte()
		if err != nil {
			break
		}

		if i >= len(problems) {
			problems = append(problems, problem{ []int{}, mul })
		}

		if char == '\n' {
			i = 0
			continue
		}

		if char == ' ' {
			if dig_i > 0 {
				i++
				dig_i = 0
				spaces = 0
			}
			spaces++
		} else if char == '+' || char == '*' {
			if char == '+' {
				problems[i].operator = sum
			}
			i++
		} else {
			dig_idx := dig_i + spaces - 1
			for dig_idx >= len(problems[i].values) {
				problems[i].values = append(problems[i].values, 0)
			}
			v := int(char) - 48 // Magic number to convert from byte to int
			problems[i].values[dig_idx] = 10*problems[i].values[dig_idx] + v
			dig_i++
		}
	}
	return problems
}

func calc_grand_total(problems *[]problem) (total int) {
	for _, p := range *problems {
		total += p.operate()
	}
	return
}

func main() {
	input, err := os.Open("example")
	if err != nil {
		panic(err)
	}

	// problems := parse_problems(input)
	problems := parse_problems_vertical(input)

	fmt.Println("Found problems:", problems)

	res := calc_grand_total(&problems)

	fmt.Println("Grand total:", res)
}
