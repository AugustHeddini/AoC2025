package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
)

type interval struct {
	start int
	end int
}

func (self *interval) String() string {
	return fmt.Sprintf("{s: %d, e: %d}", self.start, self.end)
}

type intervalReverseStart []*interval

func (a intervalReverseStart) Len() int 			{ return len(a) }
func (a intervalReverseStart) Swap(i, j int) 		{ a[i], a[j] = a[j], a[i] }
func (a intervalReverseStart) Less(i, j int) bool 	{ return a[i].start > a[j].start } 

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func contains(v int, span *interval) bool {
	return v >= span.start && v <= span.end
}

func parse_ingredient_list(input *os.File) (fresh []*interval, ingredients []int) {
	scanner := bufio.NewScanner(input)
	fresh = make([]*interval, 0)
	ingredients = make([]int, 0)

	read_intervals := true
	var start int
	var end int

	scanning:
	for scanner.Scan() {
		line := scanner.Text()

		switch read_intervals {
		case true:
			_, err := fmt.Sscanf(line, "%d-%d", &start, &end)

			if err != nil {
				read_intervals = false
				continue
			}

			for _, intr := range fresh {
				if start > intr.start && start < intr.end {
					if end > intr.end {
						intr.end = end
					}
					continue scanning
				}
			}
			fresh = append(fresh, &interval{ start, end })
		case false:
			_, err := fmt.Sscanf(line, "%d", &start)
			if err == nil {
				ingredients = append(ingredients, start)
			}
		}
	}
	return
}

func collapse_spans(fresh []*interval) []*interval {

	for idx, i := range fresh {
		for jdx := idx+1; jdx < len(fresh); jdx++ {
			j := fresh[jdx]
			if i.start >= j.start && i.start <= j.end {
				// fmt.Println("Collapsing spans", i, "and", j)
				end := max(i.end, j.end)
				collapsed_spans := append(fresh[:idx], fresh[idx+1: jdx]...)
				collapsed_spans = append(collapsed_spans, &interval{ j.start, end })
				collapsed_spans = append(collapsed_spans, fresh[jdx+1:]...)
				return collapse_spans(collapsed_spans)
			}
		}
	}
	return fresh
}

func count_fresh(fresh []*interval, ingredients []int) int {
	count := 0
	counting:
	for _, ingredient := range ingredients {
		for _, span := range fresh {
			if contains(ingredient, span) {
				count++
				continue counting
			}
		}
	}
	return count
}

func sum_intervals(spans []*interval) (count int) {
	for _, span := range spans {
		count += span.end - span.start + 1
	}
	return
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	fresh, _ := parse_ingredient_list(input)
	// fmt.Println("Found fresh: ", fresh, " and ingredients: ", ingr)
	sort.Sort(intervalReverseStart(fresh))

	fresh = collapse_spans(fresh)
	fmt.Println("Fresh:", fresh)

	// res := count_fresh(fresh, ingr)
	// fmt.Println("Found", res, "fresh ingredients")

	res := sum_intervals(fresh)
	fmt.Println("Total valid IDs:", res)
}
