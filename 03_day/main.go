package main

import (
	"fmt"
	"os"
	"bufio"
)

type battery struct {
	power int
	idx int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse_banks(input *os.File) (banks [][]int) {
	scanner := bufio.NewReader(input)

	bank := []int{} 
	for {
		byte, err := scanner.ReadByte()
		if err != nil {
			break
		}
		if byte == '\n' {
			banks = append(banks, bank)
			bank = []int{}
			continue
		}
		bank = append(bank, int(byte) - 48)
	}
	return
}

func contains(list []int, val int) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func max_val(list []int) (mx int) {
	mx = -1
	for _, v := range list {
		if v > mx {
			mx = v
		}
	}
	return
}

func max_with_skip(list []int, skip []int) (idx, mx int) {
	idx = -1
	for i, val := range list {
		if contains(skip, i) {
			continue
		}
		if val > mx {
			idx = i
			mx = val
		}
	}
	return
}

func find_joltages(banks [][]int) (sum int) {
	for _, bank := range banks {
		first := battery{ 0, -1 }		
		second := battery{ 0, -1 }		
		for i, cur_pwr := range bank {
			if cur_pwr > first.power {
				second = first
				first = battery{ cur_pwr, i }
			} else if cur_pwr > second.power {
				second = battery{ cur_pwr, i }
			}
		}
		max_joltage := 10*first.power + second.power
		if first.power == second.power || first.idx < second.idx {
			sum += max_joltage
			continue
		}

		// Case where largest digit is after second largest
		max_joltage = 10*second.power + first.power
		for i := first.idx + 1; i < len(bank); i++ {
			val := 10*first.power + bank[i]
			if val > max_joltage {
				max_joltage = val
			}
		}
		sum += max_joltage
	}
	return
}

func find_long_joltage(bank []int, rem int, curr int) int {

	if rem == 1 {
		return 10*curr + max_val(bank)
	}

	skips := []int{}
	max_idx, max_val := max_with_skip(bank, skips)
	for max_idx > len(bank) - rem {
		skips = append(skips, max_idx)
		max_idx, max_val = max_with_skip(bank, skips)
	}

	return find_long_joltage(bank[max_idx+1:], rem - 1, 10*curr + max_val)
}

func find_long_joltages(banks [][]int) (sum int) {
	for _, bank := range banks {
		sum += find_long_joltage(bank, 12, 0)
	}
	return
}

func main() {
	input, err := os.Open("input")
	check(err)
	
	banks := parse_banks(input)

	// res := find_joltages(banks)
	res := find_long_joltages(banks)

	fmt.Println("Total joltage: ", res)
}


