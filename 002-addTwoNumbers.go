package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func toArray(l *ListNode) []int {
	ln := []int{}
	for l.Next != nil {
		ln = append(ln, l.Val)
		l = l.Next
	}
	ln = append(ln, l.Val)
	return ln
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// fmt.Printf("l1: %v\n", toArray(l1))
	// fmt.Printf("l2: %v\n", toArray(l2))

	output := ListNode{0, nil}

	result := &output
	addingOne := false
	addition := 0

	for l1 != nil || l2 != nil {

		addition = 0
		if l1 != nil {
			addition += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			addition += l2.Val
			l2 = l2.Next
		}
		if addingOne {
			addition += 1
		}

		// fmt.Printf("addition: %v\n", addition)

		if addition > 9 {
			addingOne = true
			result.Val = addition % 10
		} else {
			addingOne = false
			result.Val = addition
		}
		if l1 != nil || l2 != nil {
			result.Next = &ListNode{0, nil}
			result = result.Next
		} else {
			if addingOne {
				result.Next = &ListNode{1, nil}
			}
		}

	}

	// fmt.Printf("output : %v\n", toArray(&output))
	return &output
}
