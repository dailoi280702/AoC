package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var w, h = 101, 103

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cx, cy := w/2, h/2
	quadrants := make([]int, 4)
	robots := [][]int{}
	scn := bufio.NewScanner(f)

	for scn.Scan() {
		in, err := toInts(scn.Text())
		if err != nil {
			panic(err)
		}

		px, py, vx, vy := in[0], in[1], in[2], in[3]

		x := ((vx+w)*100 + px) % w
		y := ((vy+h)*100 + py) % h

		if x < cx && y < cy {
			quadrants[0] += 1
		} else if x < cx && y > cy {
			quadrants[1] += 1
		} else if x > cx && y < cy {
			quadrants[2] += 1
		} else if x > cx && y > cy {
			quadrants[3] += 1
		}

		robots = append(robots, []int{px, py, vx, vy})
	}

	fmt.Printf("safety factor: %d\n", quadrants[0]*quadrants[1]*quadrants[2]*quadrants[3])

	for i := 0; i < 999999; i++ {
		cx, cy := centroid(robots)
		if densisty(robots, cx, cy, 1) > 0.9 {
			printBathRoom(robots)
			fmt.Printf("found a tree after %d\n seconds", i)
			break
		}

		increaseTime(robots)
	}
}

func densisty(robots [][]int, x, y, delta int) float64 {
	x -= delta
	y -= delta

	d := float64(0)

	for i := 0; i < delta*2+1; i++ {
		for j := 0; j < delta*2+1; j++ {
			for k := range robots {
				if i+x == robots[k][0] && j+y == robots[k][1] {
					d++
					break
				}
			}
		}
	}
	return d / float64((delta*2+1)*(delta*2+1))
}

func centroid(robots [][]int) (int, int) {
	x, y := 0, 0
	for i := range robots {
		x, y = x+robots[i][0], y+robots[i][1]
	}
	return x / len(robots), y / len(robots)
}

func printBathRoom(robots [][]int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ok := false
			for k := 0; k < len(robots); k++ {
				if i == robots[k][1] && j == robots[k][0] {
					ok = true
					break
				}
			}
			if ok {
				fmt.Print("##")
			} else {
				fmt.Print("..")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func increaseTime(robots [][]int) {
	for i := range robots {
		px, py, vx, vy := robots[i][0], robots[i][1], robots[i][2], robots[i][3]

		robots[i][0] = (vx%w + w + px) % w
		robots[i][1] = (vy%w + h + py) % h
	}
}

func toInts(str string) ([]int, error) {
	re := regexp.MustCompile(`-?\d+`)

	matches := re.FindAllString(str, -1)

	var integers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to integer: %w", match, err)
		}
		integers = append(integers, num)
	}

	return integers, nil
}
