package main

/**

#Approach

Declare counter
Loop over input (nums)
	if current number is different of the provided value (val variable), move the current number to the wanted position represented by nums[counter]
	Add 1 to the counter
Return counter

#Example
Given nums = [3, 2, 2, 3] and val = 3, the function will do the following:

Initialize count = 0.
Iterate through nums:
	nums[0] = 3 (equal to val), so continue.
	nums[1] = 2 (not equal to val), so set nums[count] = nums[1] (i.e., nums[0] = 2) and increment count to 1.
	nums[2] = 2 (not equal to val), so set nums[count] = nums[2] (i.e., nums[1] = 2) and increment count to 2.
	nums[3] = 3 (equal to val), so continue.

Return count = 2.
After this, nums will be [2, 2, 2, 3]. The first count = 2 elements are [2, 2],

#Complexity
	Time complexity:
		O(n)
	Space complexity:
		O(1)
*/

func removeElement(nums []int, val int) int {
	counter := 0

	for iterator := 0; iterator < len(nums); iterator++ {
		if nums[iterator] != val {
			nums[counter] = nums[iterator]
			counter++
		}
	}

	return counter
}

func main() {
	removeElement([]int{3, 2, 2, 3}, 3)             // expected return k = 2 and nums = [2,2,_,_]
	removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2) // expected return k = 5 and nums = [0,1,4,0,3,_,_,_]
}
