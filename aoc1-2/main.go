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

	var (
		sum int
		sc  string
	)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "oneight", "18")
		line = strings.ReplaceAll(line, "twone", "21")
		line = strings.ReplaceAll(line, "threeight", "38")
		line = strings.ReplaceAll(line, "fiveight", "58")
		line = strings.ReplaceAll(line, "sevenine", "79")
		line = strings.ReplaceAll(line, "eightwo", "82")
		line = strings.ReplaceAll(line, "eighthree", "83")
		line = strings.ReplaceAll(line, "nineight", "98")

		line = strings.ReplaceAll(line, "one", "1")
		line = strings.ReplaceAll(line, "two", "2")
		line = strings.ReplaceAll(line, "three", "3")
		line = strings.ReplaceAll(line, "four", "4")
		line = strings.ReplaceAll(line, "five", "5")
		line = strings.ReplaceAll(line, "six", "6")
		line = strings.ReplaceAll(line, "seven", "7")
		line = strings.ReplaceAll(line, "eight", "8")
		line = strings.ReplaceAll(line, "nine", "9")

		lastIdx := -1
		for i, c := range line {
			if '1' <= c && c <= '9' {
				sc = string(c)
				if lastIdx == -1 {
					css, _ := strconv.Atoi(sc)
					sum += 10 * css
				}
				lastIdx = i
			}
		}
		v, _ := strconv.Atoi(sc)
		sum += v
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(sum)
}
