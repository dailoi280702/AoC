package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func main() {
	f, err := os.Open("input.txt")
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

	res := 0
	for i := range graph {
		for j := range graph[0] {
			if graph[i][j] == '0' {
				res += score(graph, i, j, map[string]bool{})
			}
		}
	}

	fmt.Println(res)
}

func score(graph [][]rune, x, y int, visited map[string]bool) int {
	if len(graph) < 1 || (len(visited) == 0 && graph[x][y] != '0') {
		return 0
	}

	h := len(graph) - 1
	w := len(graph[0]) - 1
	if x < 0 || y < 0 || x > w || y > h {
		return 0
	}

	// part 1
	// visited[key(x, y)] = true
	if graph[x][y] == '9' && len(visited) != 0 {
		return 1
	}

	res := 0
	// part 2
	visited[key(x, y)] = true
	for i := range directions {
		a, b := x+directions[i][0], y+directions[i][1]
		if a < 0 || b < 0 || a > w || b > h || visited[key(a, b)] || graph[a][b]-graph[x][y] != 1 {
			continue
		}

		res += score(graph, a, b, visited)
	}

	delete(visited, key(x, y))

	return res
}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
