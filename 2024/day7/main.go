package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scn := bufio.NewScanner(f)

	var total atomic.Int64
	var wg sync.WaitGroup

	for scn.Scan() {
		l := strings.Split(strings.TrimSpace(scn.Text()), ": ")
		res := toInt(l[0])
		args := toInts(strings.Split(l[1], " "))

		wg.Add(1)
		go func(res int, args []int) {
			defer wg.Done()
			results := []int{args[0]}

			for i := 1; i < len(args); i++ {
				new_results := []int{}
				for j := range results {
					new_results = append(
						new_results,
						results[j]*args[i],
						results[j]+args[i],
						concat(results[j], args[i]),
					)
				}
				results = new_results
			}

			for i := range results {
				if results[i] == res {
					total.Add(int64(res))
					break
				}
			}
		}(res, args)
	}

	wg.Wait()

	fmt.Println(total.Load())
}

func toInts(in []string) []int {
	out := make([]int, len(in))
	for i := range in {
		out[i] = toInt(in[i])
	}
	return out
}

func toInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}

	return out
}

func concat(a, b int) int {
	c, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	if err != nil {
		panic(err)
	}

	return c
}
