package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (i Point) LessThan(j Point) bool {
	return (i.X*i.X + i.Y*i.Y) < (j.X*j.X + j.Y*j.Y)
}

func (i Point) EqualTo(j Point) bool {
	return i.X == j.X && i.Y == j.Y
}

func (p Point) String() string {
	return fmt.Sprintf("`%d,%d`", p.X, p.Y)
}

var directions = [][]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scn := bufio.NewScanner(f)

	graph := [][]rune{}
	currDirection := 0
	path := map[Point]int{}

	i := 0
	x, y := 0, 0
	for scn.Scan() {
		r := []rune{}
		for _, c := range strings.TrimSpace(scn.Text()) {
			if c == '^' {
				x, y = len(r), len(graph)
			}
			r = append(r, c)
		}
		graph = append(graph, r)
		i++
	}

	a, b := x, y

	for {
		xx := x + directions[currDirection][0]
		yy := y + directions[currDirection][1]
		if xx < 0 || yy < 0 || xx >= len(graph[0]) || yy >= len(graph) {
			break
		}
		graph[y][x] = '*'
		path[Point{X: x, Y: y}]++

		if graph[yy][xx] == '#' {
			currDirection = (currDirection + 1) % len(directions)
			continue
		}

		x = xx
		y = yy
	}
	graph[y][x] = '*'
	path[Point{X: x, Y: y}]++

	obstructions := 0
	for p := range path {
		c := graph[p.Y][p.X]
		graph[p.Y][p.X] = '#'
		if !check(graph, a, b) {
			obstructions++
		}
		graph[p.Y][p.X] = c
	}

	printMap(graph)
	fmt.Println(len(path))
	fmt.Println(obstructions)
}

func check(graph [][]rune, x, y int) bool {
	currDirection := 0
	corners := map[Point]int{}
	maxStep := int(math.Max(float64(len(graph)), float64(len(graph[0]))))

	for {

		xx := x + directions[currDirection][0]
		yy := y + directions[currDirection][1]
		if xx < 0 || yy < 0 || xx >= len(graph[0]) || yy >= len(graph) {
			return true
		}

		if graph[yy][xx] == '#' {
			p := Point{X: x, Y: y}
			corners[p]++
			if corners[p] >= maxStep {
				return false
			}

			currDirection = (currDirection + 1) % len(directions)
			continue
		}

		x = xx
		y = yy
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
}
