package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("error openning input file")
	}

	scn := bufio.NewScanner(f)

	data := [][]int{}

	for scn.Scan() {
		line := strings.Fields(scn.Text())
		row := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			row[i] = toInt(line[i])
		}
		data = append(data, row)
	}

	c := 0
	cache := map[string]float64{}

	for i := 0; i < len(data); i++ {
	inner:
		for j := 0; j <= len(data[i]); j++ {
			line := data[i]
			l, r := 0, 1
			if j == 0 {
				l = 1
				r = 2
			} else if j == 1 {
				r = 2
			}

			firstDist := line[r] - line[l]

			for r < len(line) {
				if r == j {
					r = j + 1
					continue
				}
				if l == j {
					l = j + 1
					continue
				}

				key := fmt.Sprintf("%d %d %d", i, l, r)

				dist, ok := cache[key]
				if !ok {
					dist = float64(line[r]) - float64(line[l])
					cache[key] = dist
				}

				if math.Abs(dist) < 1 || math.Abs(dist) > 3 || dist*float64(firstDist) < 0 {
					continue inner
				}

				l++
				r++
			}

			c++
			break
		}
	}

	fmt.Println(c)
}

func toInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(fmt.Sprintf("Dirty input: %+v", in))
	}

	return out
}
