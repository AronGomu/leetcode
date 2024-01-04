package main

import (
	// "fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	// fmt.Printf("s : %s\n", s)
	lenLongestSubstring := -1

	if len(s) < 2 {
		return 1
	}

	// Make loop
	i1 := 0
	i2 := 1
	nextI := 1

	for {
		// fmt.Printf("i1 : %v, i2 : %v, nextI : %v\n", i1, i2, nextI)

		if hasDuplicateChar(s[i1:i2]) || strings.Contains(s[i1:i2], string(s[nextI])) {
			i1 += 1
		}
		nextI += 1
		i2 += 1

		if nextI >= len(s) {
			break
		}
	}

	lenLongestSubstring = len(s[i1:i2])
	// fmt.Printf("[%v:%v] %v => %v\n", i1, i2, s[i1:i2], lenLongestSubstring)

	return lenLongestSubstring
}

func hasDuplicateChar(s string) bool {
	allChars := ""
	for i := range s {
		c := string(s[i])

		if strings.Contains(allChars, c) {
			return true
		}

		allChars += c
	}

	return false
}
