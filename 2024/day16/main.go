package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("example.txt")
	defer f.Close()

	scn := bufio.NewScanner(f)

	grid := [][]rune{}

	x, y := 0, 0
	for scn.Scan() {
		r := []rune{}
		for _, c := range strings.TrimSpace(scn.Text()) {
			if c == 'S' {
				x, y = len(r), len(grid)
			}
			r = append(r, c)
		}
		grid = append(grid, r)
	}

	minCost := math.MaxInt
	list := [][]int{{0, x, y, 1, 0}}
    backTrack := map[string]map[string][]int {}
    visited := map[string][]int {}
    endState := map[string][]int {}
    bestCost := map[string]int {key(list[0][1:]...): list[0][0]}

	for len(list) > 0 {
		min := 0
		for i := range list {
			if list[i][0] < list[min][0] {
				min = i
			}
		}

		p := list[min]
		if min == len(list)-1 {
			list = list[:min]
		} else {
			list = append(list[:min], list[min+1:]...)
		}

		cost, x, y, dx, dy := p[0], p[1], p[2], p[3], p[4]
        k := key(p[1:]...)
        prevBest, ok := bestCost[k];  if !ok {prevBest = math.MaxInt}
        if cost > prevBest {
            continue
        }


		if grid[y][x] == 'E' {
			if cost > minCost {
				break
			}

			minCost = cost
            endState[k] = p[1:]

            continue
		}


		for _, np := range [][]int{
			{cost + 1, x + dx, y + dy, dx, dy},
			{cost + 1000, x, y, dy, dx},
			{cost + 1000, x, y, -dy, -dx},
		} {
			// nc, nx, ny :=np[0], np[1], np[2]
			nc, nx, ny := np[0], np[1], np[2]
            nk := key(np[1:]...)
            if grid[ny][nx] == '#'  { 
                continue
            }

            prevBest, ok := bestCost[nk] 
            if !ok {
                prevBest = math.MaxInt
            }

            if nc > prevBest {
                continue
            }

            if nc < prevBest {
                bestCost[nk] = nc
                delete(backTrack, nk)
            }

            if _, ok := backTrack[nk]; !ok {
                backTrack[nk] = map[string][]int{}
            }

            backTrack[nk][k] = p[1:]

            list = append(list, np)
		}
	}

    queue := []string {}
    for k  := range endState {
        queue = append(queue, k)
    }

    for len(queue) > 0 {
        key := queue[0]
        queue = queue[1:]

        for last, p := range backTrack[key] {
            visited[last] = p

            queue = append(queue, last)
        }
    }

    bestSpot := map[string][]int{}
    for _, p := range visited {
        x, y := p[0], p[1]
        bestSpot[key(x, y)] = p[:2]
    }

    for _,p := range bestSpot {
        grid[p[1]][p[0]] = 'O'
    }
    

    printMap(grid)
    fmt.Println(bestSpot)
	fmt.Println(minCost+1)
}

func printMap(grid [][]rune) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Print(string(grid[i][j]))
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func key(in ...int) string {
	return fmt.Sprintf("%v", in)
}
