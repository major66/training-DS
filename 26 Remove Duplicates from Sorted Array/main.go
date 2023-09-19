package main

/*
*
Approach

  - Check if array is empty and return 0 if it's the case.

  - Initializes a counter variable uniqueCount to 1

  - Iterate through the nums slice starting from the second element.

  - If the current element is different from the previous one.

  - move the element to the uniqueCount-th position in nums.

  - increments uniqueCount

  - Return uniqueCount, which is the number of unique elements in nums.

    This function modifies the nums slice in-place and returns the number of unique elements.

Example

	If the input is [1, 1, 2, 2, 2, 2], the function removeDuplicates will work as follows:

	1. Initialize uniqueCount to 1 because the array is not empty.
	2. Start iterating from the second element (index 1).
		- At index 1: nums[1] is 1, which is the same as nums[0]. So, continue to the next iteration.
		- At index 2: nums[2] is 2, which is different from nums[1]. Move nums[2] to nums[uniqueCount] (nums[1] = 2) and increment uniqueCount to 2.
		- At index 3, 4, and 5: nums[3], nums[4], and nums[5] are all 2, which are the same as nums[2]. So, continue to the next iteration.
	3. Return uniqueCount, which is 2.
	The array nums will be modified to [1, 2, 2, 2, 2, 2]. The first uniqueCount = 2 elements [1, 2] are unique and in the order they were initially present in nums. The remaining elements [2, 2, 2, 2] are not important, as specified in the problem statement.

	So, the output will be 2, and nums will be effectively [1, 2, _, _, _, _], where _ denotes elements that are not important.

Complexity

	Time complexity:
		O(n)
	Space complexity:
		O(1)
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	uniqueCount := 1

	for iterator := 1; iterator < len(nums); iterator++ {
		if nums[iterator] != nums[iterator-1] {
			nums[uniqueCount] = nums[iterator]
			uniqueCount++
		}
	}

	return uniqueCount
}

func main() {
	removeDuplicates([]int{1, 1, 2})
	removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
}
