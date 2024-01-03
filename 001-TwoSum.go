package main

func main() {
	twoSum([]int{3, 2, 4}, 6)
}

func twoSum(nums []int, target int) []int {
	a := 0
	b := 1
	found := false

	for found != true {
		for i := b; i < len(nums); i++ {
			if nums[a]+nums[i] == target {
				return []int{a, i}
			}
		}
		a += 1
		b += 1

		if b >= len(nums) {
			found = true
		}
	}

	return []int{0, 0}
}
