func canJump(nums []int) bool {
	// what is the farthest index we can reach
	maxReach := 0
	n := len(nums)

	// iterate over the numbers
	for i := range n {
		// if the curr index surpasses the maxreach we return false
		if i > maxReach {
			return false
		}
		
		if i +nums[i] > maxReach {
			maxReach = i+nums[i]
		}

	// if we can jump
		if maxReach >= n-1 {
			return true
		}
	}
	// if it completes the loop we return true
	return true
    
}
