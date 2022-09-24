package Solutions

func findUnsortedSubarray(nums []int) int {
	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))
	dpMax[0] = 0
	dpMin[len(nums)-1] = len(nums)-1
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[dpMax[i-1]] {
			dpMax[i] = i
		} else {
			dpMax[i] = dpMax[i-1]
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= nums[dpMin[i+1]] {
			dpMin[i] = i
		} else {
			dpMin[i] = dpMin[i+1]
		}
	}
	var left = -1
	for i := 0; i < len(nums); i++ {
		if dpMax[i]!=dpMin[i]{
			left=i
            break
		}
	}
	var right=-1
	for i:=len(nums)-1;i>=0;i--{
		if dpMin[i]!=dpMax[i]{
			right=i
            break
		}
	}
  
    if left==-1&&right ==-1{
        return 0
    }
	return right-left+1
}