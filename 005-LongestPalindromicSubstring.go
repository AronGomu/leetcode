package main

import "fmt"

func longestPalindrome(s string) string {

	if len(s) == 1 {
		return s
	}

	// fmt.Printf("s: %s\n", s)
	pal := s[0:1]

	p1 := 0
	nextPalLength := 2
	p2 := p1 + nextPalLength

	for {
		// fmt.Printf("%v > %v\n", nextPalLength, len(s))
		if nextPalLength > len(s) {
			break
		}

		// fmt.Printf("s[%v:%v] => %s\n", p1, p2, s[p1:p2])
		if !isPalindromeReverseString(s[p1:p2]) {
			p1 += 1
			p2 += 1
			if p2 > len(s) {
				nextPalLength += 1
				p1 = 0
				p2 = p1 + nextPalLength
			}
		} else {
			pal = s[p1:p2]
			nextPalLength += 1
			p1 = 0
			p2 = p1 + nextPalLength
		}

	}

	fmt.Printf("output: %s\n", pal)
	return pal
}

func isPalindromeReverseString(s string) bool {
	// fmt.Printf("isPalindromeReverseString substring: %s\n", s)
	var reverseS string = ""
	var length = len(s)

	for i := length - 1; i >= 0; i-- {
		reverseS = reverseS + string(s[i])
	}

	if s == reverseS {
		// fmt.Println("isPalindromeReverseString TRUE")
		return true
	}

	// fmt.Println("isPalindromeReverseString FALSE")
	return false
}
