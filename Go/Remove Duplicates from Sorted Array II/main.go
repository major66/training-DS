package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	writeIndex := 0
	readIndex := 0
	freq := 0

	for readIndex < len(nums) {
		if readIndex == 0 || nums[readIndex] != nums[readIndex-1] {
			freq = 1
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		} else {
			freq++
			if freq <= 2 {
				nums[writeIndex] = nums[readIndex]
				writeIndex++
			}
		}
		readIndex++
	}

	return writeIndex
}

func main() {
	removeDuplicates([]int{1, 1, 1, 2, 2, 3}) // expected [1, 1, 2, 2, 3]
	removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})
}
