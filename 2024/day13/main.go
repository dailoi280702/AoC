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
		panic(err)
	}

	total := 0

	scn := bufio.NewScanner(f)
	for scn.Scan() {
		l := strings.Split(strings.TrimSpace(scn.Text()), " ")
		ax, ay := toInt(l[2][2:len(l[2])-1]), toInt(l[3][2:])

		scn.Scan()
		l = strings.Split(strings.TrimSpace(scn.Text()), " ")
		bx, by := toInt(l[2][2:len(l[2])-1]), toInt(l[3][2:])

		scn.Scan()
		l = strings.Split(strings.TrimSpace(scn.Text()), " ")
		px, py := toInt(l[1][2:len(l[1])-1]), toInt(l[2][2:])

		px += 10000000000000
		py += 10000000000000

		i := float64(px*by-py*bx) / float64(ax*by-ay*bx)
		j := float64(float64(px)-float64(ax)*i) / float64(bx)

		if isInt(i) && isInt(j) {
			total += int(i*3 + j)
		}

		scn.Scan()
	}

	fmt.Println(total)
}

func toInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}

	return out
}

func isInt(num float64) bool {
	return math.Abs(num-float64(int(num))) < 1e-9
}
