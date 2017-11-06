package main

// store complement of ``nums`` to ``target``: remains[target-nums[n]] = n
// solution got when complement remains[target-n] found
func twoSum(nums []int, target int) []int {
	remains := make(map[int]int, len(nums))
	for i, n := range nums {
		if index, found := remains[n]; found {
			return []int{index, i}
		}

		remains[target-n] = i
	}

	return nil
}
