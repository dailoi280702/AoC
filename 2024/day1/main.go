package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error openning input file: %v", err)
	}
	defer f.Close()

	left := []int{}
	right := []int{}

	scn := bufio.NewScanner(f)
	for scn.Scan() {
		inputs := strings.Fields(scn.Text())
		if len(inputs) != 2 {
			panic(fmt.Sprintf("Dirty input: %+v", inputs))
		}

		left = append(left, strToNum(inputs[0]))
		right = append(right, strToNum(inputs[1]))
	}

	slices.Sort(left)
	slices.Sort(right)

	distance := 0
	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i]) - float64(right[i])))
	}

	fmt.Println(distance)

	similarities := 0
	count := map[int]int{}
	for i := 0; i < len(right); i++ {
		c, ok := count[right[i]]
		if ok {
			count[right[i]] = c + 1
		} else {
			count[right[i]] = 1
		}
	}

	for i := 0; i < len(left); i++ {
		c, ok := count[left[i]]
		if ok {
			similarities += left[i] * c
		}
	}

	fmt.Println(similarities)
}

func strToNum(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(fmt.Sprintf("Dirty input: %+v", in))
	}

	return out
}
