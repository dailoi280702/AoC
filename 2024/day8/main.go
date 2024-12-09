package main

import (
	"bufio"
	"fmt"
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

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scn := bufio.NewScanner(f)

	antena := map[rune]map[Point]bool{}
	graph := [][]rune{}
	antinode := map[Point]bool{}

	i := 0
	for scn.Scan() {
		r := []rune{}
		for j, c := range strings.TrimSpace(scn.Text()) {
			r = append(r, c)
			if c != '.' {
				p := Point{X: i, Y: j}
				if _, ok := antena[c]; ok {
					antena[c][p] = true
				} else {
					antena[c] = map[Point]bool{p: true}
				}
			}
		}
		graph = append(graph, r)
		i++
	}

	for _, loc := range antena {
		for p1 := range loc {
			for p2 := range loc {
				if p1 == p2 {
					continue
				}

				antinode[p1] = true
				antinode[p2] = true

				for _, p := range findPointsInRectangle(p1, p2, len(graph[0])-1, len(graph)-1) {
					antinode[p] = true
					if graph[p.X][p.Y] == '.' {
						graph[p.X][p.Y] = '#'
					}
				}
			}
		}
	}

	printMap(graph)
	fmt.Println(len(antinode))
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

func findPointsInRectangle(p1 Point, p2 Point, w int, l int) []Point {
	points := []Point{}

	if p1 == p2 {
		return points
	}

	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	x, y := p1.X, p1.Y
	for {
		x -= dx
		y -= dy
		if x < 0 || x > w || y < 0 || y > l {
			break
		}

		points = append(points, Point{X: x, Y: y})
	}

	x, y = p2.X, p2.Y
	for {
		x += dx
		y += dy
		if x < 0 || x > w || y < 0 || y > l {
			break
		}

		points = append(points, Point{X: x, Y: y})
	}

	return points
}
