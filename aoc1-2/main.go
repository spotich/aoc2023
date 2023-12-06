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

	var sum int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line2 := line
		fmt.Println(line)
		line = strings.ReplaceAll(line, "eight", "8")
		line = strings.ReplaceAll(line, "seven", "7")
		line = strings.ReplaceAll(line, "one", "1")
		line = strings.ReplaceAll(line, "two", "2")
		line = strings.ReplaceAll(line, "three", "3")
		line = strings.ReplaceAll(line, "four", "4")
		line = strings.ReplaceAll(line, "five", "5")
		line = strings.ReplaceAll(line, "six", "6")
		line = strings.ReplaceAll(line, "nine", "9")
		line2 = strings.ReplaceAll(line2, "seven", "7")
		line2 = strings.ReplaceAll(line2, "nine", "9")
		line2 = strings.ReplaceAll(line2, "one", "1")
		line2 = strings.ReplaceAll(line2, "eighthree", "83")
		line2 = strings.ReplaceAll(line2, "eight", "8")
		line2 = strings.ReplaceAll(line2, "two", "2")
		line2 = strings.ReplaceAll(line2, "three", "3")
		line2 = strings.ReplaceAll(line2, "four", "4")
		line2 = strings.ReplaceAll(line2, "five", "5")
		line2 = strings.ReplaceAll(line2, "six", "6")

		lastIdx := -1
		number := 0
		for i, c := range line {
			if '1' <= c && c <= '9' {
				lastIdx = i
			}
		}
		cs := string(line[lastIdx])

		lastIdx = -1
		for _, c := range line2 {
			if '1' <= c && c <= '9' {
				cs := string(c)
				if lastIdx == -1 {
					css, _ := strconv.Atoi(cs)
					number += 10 * css
				}
			}
		}
		v, _ := strconv.Atoi(cs)
		number += v
		fmt.Println(number)
		sum += number
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(sum)
}

// func find(line string) int {
// 	m := map[string]int{
// 		"one":   1,
// 		"two":   2,
// 		"three": 3,
// 		"four":  4,
// 		"five":  5,
// 		"six":   6,
// 		"seven": 7,
// 		"eight": 8,
// 		"nine":  9,
// 	}
// 	for i, c := range line {
// 		lastIdx := -1
// 		firstIdx := -1
// 		if '1' <= c && c <= '9' {
// 			cs := string(c)
// 			if firstIdx == -1 {
// 				firstIdx := i
// 			}
// 			lastIdx = i
// 		}

// 	}
// 	firstsubstrin =
// 	lastsubstring =
// 	lastsubstring = strings.LastIndex(lastsubstring, "seven")
// 	lastsubstring = strings.ReplaceAll(lastsubstring, "seven", "7")
// 	lastsubstring = strings.ReplaceAll(lastsubstring, "one", "1")
// 	lastsubstring = strings.ReplaceAll(lastsubstring, "two", "2")
// 	lastsubstring = strings.ReplaceAll(lastsubstring, "three", "3")
// 	lastsubstring = strings.ReplaceAll(lastsubstring, "four", "4")
// 	lastsubstring = strings.ReplaceAll(lastsubstring, "five", "5")
// 	lastsubstring = strings.ReplaceAll(lastsubstring, "six", "6")
// 	lastsubstring = strings.ReplaceAll(lastsubstring, "nine", "9")
// 	for k,v := range m {

// 	}

// }
