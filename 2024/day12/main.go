package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// counter clockwise directions
var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	scn := bufio.NewScanner(f)
	graph := [][]rune{}

	for scn.Scan() {
		l := strings.TrimSpace(scn.Text())
		graph = append(graph, make([]rune, len(l)))

		for i, r := range l {
			graph[len(graph)-1][i] = r
		}
	}

	price := 0
	discountedPrice := 0
	visited := map[string]bool{}
	for i := range graph {
		for j := range graph[0] {
			if !visited[key(i, j)] {
				area, piremeter, side := areaAndPerimeter(graph, i, j, visited)
				price += area * piremeter
				discountedPrice += area * side
				fmt.Printf("%c: %d, %d, %d\n", graph[i][j], area, piremeter, side)
			}
		}
	}

	fmt.Println(price)
	fmt.Println(discountedPrice)
}

func areaAndPerimeter(graph [][]rune, x, y int, visited map[string]bool) (int, int, int) {
	area, piremeter, side := 1, 0, corner(graph, x, y)

	h := len(graph) - 1
	w := len(graph[0]) - 1
	// if x < 0 || y < 0 || x > w || y > h {
	// 	return 0, 1, currSide
	// }

	visited[key(x, y)] = true
	for _, v := range directions {
		a, b := x+v[0], y+v[1]
		if a < 0 || b < 0 || a > h || b > w || graph[a][b]-graph[x][y] != 0 {
			piremeter++
			continue
		}

		if !visited[key(a, b)] {
			a, p, s := areaAndPerimeter(graph, a, b, visited)
			area += a
			piremeter += p
			side += s
		}
	}

	return area, piremeter, side
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

// B is a corner if B belongs to the following cases:
// BA  BB  BA
// AB  BA  AA
func corner(graph [][]rune, x, y int) int {
	res := 0
	for i, d1 := range directions {
		d2 := directions[(i+1)%len(directions)]
		if (equal(graph, x, y, x+d1[0], y+d1[1]) &&
			equal(graph, x, y, x+d2[0], y+d2[1]) &&
			!equal(graph, x, y, x+d1[0]+d2[0], y+d1[1]+d2[1])) ||
			(!equal(graph, x, y, x+d1[0], y+d1[1]) &&
				!equal(graph, x, y, x+d2[0], y+d2[1]) &&
				equal(graph, x, y, x+d1[0]+d2[0], y+d1[1]+d2[1])) ||
			(!equal(graph, x, y, x+d1[0], y+d1[1]) &&
				!equal(graph, x, y, x+d2[0], y+d2[1]) &&
				!equal(graph, x, y, x+d1[0]+d2[0], y+d1[1]+d2[1])) {
			res++
		}
	}

	return res
}

func equal(graph [][]rune, x, y, i, j int) bool {
	return true &&
		x >= 0 && y >= 0 &&
		i >= 0 && j >= 0 &&
		x < len(graph) && y < len(graph[0]) &&
		i < len(graph) && j < len(graph[0]) &&
		graph[x][y] == graph[i][j]
}
