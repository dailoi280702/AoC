package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	id := 0
	disk := []int{}
	mFile := map[int]int{}
	mSpace := map[int]int{}

	for {
		if files, ok := readRune(reader); ok {
			for i := 0; i < toInt(string(files)); i++ {
				disk = append(disk, id)
			}
			mFile[id] = toInt(string(files))
		} else {
			id--
			break
		}

		if spaces, ok := readRune(reader); ok {
			for i := 0; i < toInt(string(spaces)); i++ {
				disk = append(disk, -1)
			}
			mSpace[id] = toInt(string(spaces))
		} else {
			break
		}

		id++
	}

	// Part 1
	// i, j := 0, len(disk)-1
	// for i != j && i < len(disk) && j > 0 {
	// 	for disk[i] != -1 && i < len(disk)-1 {
	// 		i++
	// 	}
	//
	// 	for disk[j] == -1 && j > 1 {
	// 		j--
	// 	}
	//
	// 	if j <= i {
	// 		break
	// 	}
	//
	// 	disk[i] = disk[j]
	// 	disk[j] = -1
	//
	// 	i++
	// 	j--
	// }

	// Part 2
	i, j := 0, len(disk)-1
	for j > 0 {
		for disk[j] == -1 && j > 1 {
			j--
		}

		m := j - 1
		for disk[j] == disk[m] && m > 1 {
			m--
		}

		chunkSize := j - m

		i = 0
		for i < len(disk)-1 && i <= m {
			for disk[i] != -1 && i < len(disk)-1 {
				i++
			}

			if i > m {
				break
			}

			n := i + 1
			for disk[n] == -1 && n < len(disk)-1 {
				n++
			}

			spaceSize := n - i

			// println(fmt.Sprintf("%d, %d, %d", disk[j], chunkSize, spaceSize))
			if chunkSize > spaceSize {
				i = n
			} else {
				for j > m {
					disk[i] = disk[j]
					disk[j] = -1
					j--
					i++
				}
				break
			}
		}

		j = m
	}

	checksum := 0
	for i := range disk {
		if disk[i] == -1 {
			continue
		}

		checksum += i * disk[i]
	}
	fmt.Println(checksum)
}

func readRune(reader *bufio.Reader) (rune, bool) {
	r, _, err := reader.ReadRune()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return r, false
		}

		panic(err)
	}

	return r, unicode.IsDigit(r)
}

func toInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}

	return out
}
