package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var movement = map[rune][]int{
	'^': {0, -1},
	'v': {0, 1},
	'<': {-1, 0},
	'>': {1, 0},
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scn := bufio.NewScanner(f)
	graph := [][]rune{}
	movements := []rune{}
	px, py := 0, 0

	for scn.Scan() {
		l := strings.TrimSpace(scn.Text())
		if len(l) == 0 {
			break
		}

		row := make([]rune, 0)
		for i, r := range l {
			switch r {
			case '#':
				row = append(row, '#', '#')
			case '.':
				row = append(row, '.', '.')
			case 'O':
				row = append(row, '[', ']')
			case '@':
				row = append(row, '@', '.')
				px, py = i*2, len(graph)
			default:
				panic(fmt.Sprintf("Invalid character %c", r))
			}
		}

		graph = append(graph, row)
	}

	for scn.Scan() {
		l := strings.TrimSpace(scn.Text())
		for _, r := range l {
			movements = append(movements, r)
		}
	}
	// printMap(graph)

	for _, r := range movements {
		dx, dy := movement[r][0], movement[r][1]
		stack := pushStack(graph, px+dx, py+dy, dx, dy)
		for i := len(stack) - 1; i >= 0; i-- {
			bl, br, by := stack[i][0], stack[i][1], stack[i][2]
			graph[by][bl] = '.'
			graph[by][br] = '.'
			graph[by+dy][bl+dx] = '['
			graph[by+dy][br+dx] = ']'
		}

		if graph[py+dy][px+dx] == '.' {
			graph[py][px] = '.'
			px += dx
			py += dy
			graph[py][px] = '@'

			// fmt.Printf("move %c\n", r)
			// printMap(graph)
		}
	}

	coordinates := 0
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] == '[' {
				coordinates += 100*i + j
			}
		}
	}

	fmt.Println(coordinates)
}

func pushStack(graph [][]rune, x, y, dx, dy int) [][]int {
	stack := [][]int{}
	tracker := map[string]bool{}

	if dy == 0 {
		for x > 0 && x < len(graph[0])-1 {
			switch graph[y][x] {
			case '.':
				return stack
			case '#':
				return [][]int{}
			}

			l, r, _, ok := box(graph, x, y)
			if !ok {
				panic("Opps..")
			}

			if !tracker[key(l, r, y)] {
				stack = append(stack, []int{l, r, y})
				tracker[key(l, r, y)] = true
			}
			x += dx
		}

		return stack
	}

	if dx == 0 {
		l, r, _, ok := box(graph, x, y)
		if !ok {
			return [][]int{}
		}

		stack = append(stack, []int{l, r, y})
		tracker[key(l, r, y)] = true
		y += dy

		for y >= 0 && y < len(graph) {
			for _, b := range stack {
				bl, br, by := b[0], b[1], b[2]

				if by == y-dy {
					if graph[y][br] == '#' || graph[y][bl] == '#' {
						return [][]int{}
					}

					if l, r, _, ok := box(graph, bl, y); ok && !tracker[key(l, r, y)] {
						stack = append(stack, []int{l, r, by + dy})
						tracker[key(l, r, y)] = true
					}

					if l, r, _, ok := box(graph, br, y); ok && !tracker[key(l, r, y)] {
						stack = append(stack, []int{l, r, by + dy})
						tracker[key(l, r, y)] = true
					}
				}
			}

			y += dy
		}

		return stack
	}

	panic("Invalid direction")
}

func box(graph [][]rune, x, y int) (int, int, int, bool) {
	switch graph[y][x] {
	case '[':
		return x, x + 1, y, true
	case ']':
		return x - 1, x, y, true
	default:
		return 0, 0, 0, false
	}
}

func printMap(graph [][]rune) {
	for i := range graph {
		for j := range graph[i] {
			fmt.Print(string(graph[i][j]))
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func key(x, y, z int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, z)
}
