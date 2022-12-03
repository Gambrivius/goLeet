package goLeet

func twoSum(nums []int, target int) []int {
	// naive O(n^2) implementation beats 8.96% on runtime and 65.9% on memory
	/*
	for i := range nums {
		for j := range nums {
			if i!=j {
				if nums[i] + nums[j] == target {
					return []int{i, j}
				}
			}
		}
	}*/

	// final slight optimization, beats 31.7% on runtime and 93.83% on memory
	// avoids checking the same indices twice
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i+1; j < l; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
    return nil
}

