package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	dict := map[int]bool{}
	rules := [][]int{}
	updates := [][]int{}

	scn := bufio.NewScanner(f)
	for scn.Scan() {
		l := strings.TrimSpace(scn.Text())
		if len(l) == 0 {
			break
		}

		pages := toInts(strings.Split(l, "|"))
		if len(pages) != 2 {
			panic(fmt.Sprintf("unexpected line: `%s`", l))
		}

		dict[pages[0]] = true
		dict[pages[1]] = true
		rules = append(rules, []int{pages[0], pages[1]})
	}

	for scn.Scan() {
		l := strings.TrimSpace(scn.Text())
		pages := toInts(strings.Split(l, ","))
		updates = append(updates, pages)
	}

	res := 0
	res2 := 0

	for _, update := range updates {
		if ok, mid := valid(rules, update); ok {
			res += mid
		} else {
			fixed := fix(rules, update)
			res2 += fixed[len(fixed)/2]
		}
	}

	fmt.Println(res)
	fmt.Println(res2)
}

func valid(rules [][]int, update []int) (bool, int) {
	idx := map[int]int{}

	for i, page := range update {
		idx[page] = i + 1
	}

	for _, rule := range rules {
		if idx[rule[0]] != 0 && idx[rule[1]] != 0 && idx[rule[0]] > idx[rule[1]] {
			return false, 0
		}
	}

	return true, update[len(update)/2]
}

func fix(rules [][]int, update []int) []int {
	graph := map[int][]int{}
	indegree := map[int]int{}
	updateDict := map[int]bool{}

	for _, page := range update {
		updateDict[page] = true
	}

	for _, rule := range rules {
		if !updateDict[rule[0]] || !updateDict[rule[1]] {
			continue
		}

		if node, ok := graph[rule[0]]; ok {
			graph[rule[0]] = append(node, rule[1])
		} else {
			graph[rule[0]] = []int{rule[1]}
		}

		if degree, ok := indegree[rule[1]]; ok {
			indegree[rule[1]] = degree + 1
		} else {
			indegree[rule[1]] = 1
		}

		if _, ok := indegree[rule[0]]; !ok {
			indegree[rule[0]] = 0
		}

	}

	min := update[0]
	for _, degree := range indegree {
		if degree < min {
			min = degree
		}
	}

	queue := []int{}
	for _, page := range update {
		if indegree[page] == min {
			queue = append(queue, page)
		}
	}

	if len(queue) == 0 {
		panic("some ting went wrong!")
	}

	fixed := []int{}

	for len(fixed) < len(update) {
		node := queue[0]
		queue = queue[1:]
		fixed = append(fixed, node)

		for _, neighbor := range graph[node] {
			indegree[neighbor] -= 1
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return fixed
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
