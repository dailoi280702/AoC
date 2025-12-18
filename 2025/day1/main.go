package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scn := bufio.NewScanner(f)
	pos := 50
	pass := 0

	for scn.Scan() {
		l := scn.Text()
		direction := l[0]
		steps, _ := strconv.Atoi(l[1:])

		for range steps {
			if direction == 'R' {
				pos = (pos + 1) % 100
			} else {
				pos = (pos + 99) % 100
			}

			if pos == 0 {
				pass += 1
			}
		}
	}

	fmt.Println(pass)
}
