package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // To handle cases where k is greater than the length of the array

	reverse(nums, 0, n-1) // Reverse the entire array
	reverse(nums, 0, k-1) // Reverse the first k elements
	reverse(nums, k, n-1) // Reverse the remaining elements
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start] // Swap elements
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums) // Output: [5 6 7 1 2 3 4]
}
