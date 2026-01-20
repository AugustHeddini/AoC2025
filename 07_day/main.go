package main

import(
	"fmt"
	"bufio"
	"os"
)

type coord struct {
	x int
	y int
}

func (a *coord) move() {
	a.y++
}

func parse_manifold(input *os.File) (start coord, splitters map[coord]bool, dims coord) {
	start = coord{ -1, -1 }
	splitters = make(map[coord]bool)
	reader := bufio.NewReader(input)

	y := 0
	x := 0
	for {
		char, err := reader.ReadByte()
		if err != nil {
			break
		}

		if char == '\n' {
			x = 0
			y++
			continue
		}

		if char == 'S' {
			start = coord{ x, y }
		}

		if char == '^' {
			splitters[coord{ x, y }] = true
		}

		x++
	}

	dims = coord{ x, y }
	return
}

func propagate_tachyons(start coord, splitters map[coord]bool, dims coord) (splits int) {

	beams := []coord{ start }
	visited := make(map[coord]bool)
	for len(beams) != 0 {
		next_beams := []coord{}
		for _, beam := range beams {
			beam.move()
			if beam.y >= dims.y {
				continue
			}
			if splitters[beam] {
				left := coord{ beam.x - 1, beam.y }
				right := coord{ beam.x + 1, beam.y }
				if !visited[left] {
					next_beams = append(next_beams, left)
					visited[left] = true
				}
				if !visited[right] {
					next_beams = append(next_beams, right)
					visited[right] = true
				}
				splits++
			} else {
				if !visited[beam] {
					next_beams = append(next_beams, beam)
					visited[beam] = true
				}
			}
		}
		beams = next_beams
	}

	return
}

func count_paths(start coord, splitters map[coord]bool, dims coord) (sum int) {

	bottoms := make(map[coord]int)
	beam_count := map[coord]int{ start: 1 }
	
	for len(beam_count) != 0 {
		new_beam_count := make(map[coord]int)
		for beam, count := range beam_count {
			beam.move()
			if beam.y == dims.y {
				bottoms[beam] += count
				continue
			}

			if splitters[beam] {
				new_beam_count[coord{ beam.x - 1, beam.y }] += count
				new_beam_count[coord{ beam.x + 1, beam.y }] += count
			} else {
				new_beam_count[beam] += count
			}
		}
		beam_count = new_beam_count
	}

	for _, count := range bottoms {
		sum += count
	}

	return
}

func main() {

	input, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	start, splitters, dims := parse_manifold(input)

	// res := propagate_tachyons(start, splitters, dims)
	res := count_paths(start, splitters, dims)

	fmt.Println(res)

}
