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

	var sum int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lastIdx := -1
		number := 0
		for i, c := range line {
			if '0' <= c && c <= '9' {
				cs := string(c)
				if lastIdx == -1 {
					css, _ := strconv.Atoi(cs)
					number += 10 * css
				}
				lastIdx = i
			}
		}
		v, _ := strconv.Atoi(string(line[lastIdx]))
		number += v
		fmt.Println(number)
		sum += number
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(sum)
}
