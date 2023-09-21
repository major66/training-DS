package main

func majorityElement(nums []int) int {
	candidate, count := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
			count = 1
		} else if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func main() {
	majorityElement([]int{3, 2, 3})
}
