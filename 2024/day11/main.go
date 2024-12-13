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
    if err   != nil {
        panic(err)
    }

    scn := bufio.NewScanner(f)
    if !scn.Scan() {
        panic("Opps...")
    }

    nums := map[int]int{}
    for _, num := range strings.Split(strings.TrimSpace(scn.Text()), " ") {
        nums[toInt(num)]++
    }

    for i := 0; i < 75; i ++ {
        nums2 := map[int]int{}
        for num, count := range nums {
            for _, v := range blink(num) {
                nums2[v] += count
            }
        }
        nums = nums2
    }

    res := 0
    for num, count := range nums {

        if count < 0 {
            panic(fmt.Sprintf("Opps.. num: %d, count: %d", num, count))
        }
        res += count
    }

    fmt.Println(res)
}

func blink(num int) []int {
    if num == 0 {
        return []int{1}
    }

    str := strconv.Itoa(num)
    if len(str) % 2 == 0 {
        half := len(str) / 2
        return []int{toInt(str[:half]), toInt(str[half:])}
    }

    return []int{num * 2024}
}

func toInt(in string) int {
    out, err := strconv.Atoi(in)
    if err != nil {
        panic(err)
    }

    return out
}
