package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

const tabChar = "\t"

var (
	lines         = []string{}
	minPrefixLen  = math.MaxInt64
	tabCharSpaces = strings.Repeat(" ", 2)
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		l := strings.ReplaceAll(sc.Text(), tabChar, tabCharSpaces)
		lines = append(lines, l)
	}

	for _, l := range lines {
		if v := countLeadingSpaces(l); v < minPrefixLen {
			minPrefixLen = v
		}
	}

	for _, l := range lines {
		fmt.Println(string([]rune(l)[minPrefixLen:]))
	}
}

func countLeadingSpaces(s string) int {
	var count int
	for _, c := range s {
		if !unicode.IsSpace(c) {
			break
		}
		count++
	}
	return count
}
