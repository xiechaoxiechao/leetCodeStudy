package Solutions

func prevPermOpt1(arr []int) []int {
	p := -1
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			p = i - 1
		}
	}
	if p == -1 {
		return arr
	}
	l, r := p+1, len(arr)-1
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] >= arr[p] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	arr[p], arr[r] = arr[r], arr[p]
	return arr
}
