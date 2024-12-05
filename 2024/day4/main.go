package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scn := bufio.NewScanner(f)
	input := [][]string{}

	for scn.Scan() {
		input = append(input, strings.Split(strings.TrimSpace(scn.Text()), ""))
	}

	res := map[string]bool{}
	for i := range input {
		for j := range input[i] {
			search(input, i, j, res)
		}
	}

	count := 0
	for i := 1; i < len(input[0])-1; i++ {
		for j := 1; j < len(input)-1; j++ {
			if input[i][j] != "A" ||
				!(fmt.Sprintf("%s%s", input[i-1][j-1], input[i+1][j+1]) == "MS" ||
					fmt.Sprintf("%s%s", input[i-1][j-1], input[i+1][j+1]) == "SM") ||
				!(fmt.Sprintf("%s%s", input[i-1][j+1], input[i+1][j-1]) == "MS" ||
					fmt.Sprintf("%s%s", input[i-1][j+1], input[i+1][j-1]) == "SM") {
				continue
			}
			count++
		}
	}

	fmt.Println(len(res))
	fmt.Println(count)
}

var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 1},
	{1, 0},
	{1, -1},
	{-1, 1},
	{-1, -1},
	{-1, 0},
}

const (
	SEARCH_STRING          = "XMAS"
	REVERSED_SEARCH_STRING = "SAMX"
)

func search(input [][]string, x int, y int, foundLoc map[string]bool) {
outer:
	for i := range directions {
		endX := x + directions[i][0]*(len(SEARCH_STRING)-1)
		endY := y + directions[i][1]*(len(SEARCH_STRING)-1)
		if endX < 0 || endX >= len(input[0]) || endY < 0 || endY >= len(input) {
			continue
		}

		key := fmt.Sprintf("%d%d%d%d", x, y, endX, endY)
		if _, ok := foundLoc[key]; ok {
			continue
		}

		seq := input[x][y]
		for j := 1; j < len(SEARCH_STRING); j++ {
			seq += input[x+j*directions[i][0]][y+j*directions[i][1]]

			if !strings.Contains(SEARCH_STRING, seq) && !strings.Contains(REVERSED_SEARCH_STRING, seq) {
				continue outer
			}
		}

		if seq == SEARCH_STRING || seq == REVERSED_SEARCH_STRING {
			key := fmt.Sprintf("%d%d%d%d", endX, endY, x, y)
			foundLoc[key] = true
		}
	}
}
