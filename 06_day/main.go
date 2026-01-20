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
	if len(l) == 0 {
		return 0
	}

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

func all(l []byte, b byte) bool {
	for _, v := range l {
		if v != b {
			return false
		}
	}
	return true
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
	transpose := make([][]byte, 0)
	for {
		char, err := reader.ReadByte()
		if err != nil {
			break
		}

		if char == '\n' {
			i = 0
			continue
		}

		if i >= len(transpose) {
			transpose = append(transpose, []byte{})
		}

		transpose[i] = append(transpose[i], char)
		i++
	}

	prob := problem{ []int{}, sum }
	for _, col := range transpose {
		if all(col, ' ') {
			problems = append(problems, prob)
			prob = problem{ []int{}, sum }
			continue
		}

		val := 0
		for _, char := range col {
			if char == ' ' || char == '+' {
				continue
			} else if char == '*' {
				prob.operator = mul
			} else {
				val = 10*val + (int(char) - 48)
			}
		}
		prob.values = append(prob.values, val)
	}
	problems = append(problems, prob)

	return problems
}

func calc_grand_total(problems *[]problem) (total int) {
	for _, p := range *problems {
		total += p.operate()
	}
	return
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	// problems := parse_problems(input)
	problems := parse_problems_vertical(input)

	res := calc_grand_total(&problems)

	fmt.Println("Grand total:", res)
}
