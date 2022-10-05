package Solutions

func isPalindrome(head *ListNode) bool {
	var arr = make([]int, 0, 32)
	for i := head; i != nil; i = i.Next {
		arr = append(arr, i.Val)
	}

	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		if arr[l] != arr[r] {
			return false
		}

	}
	return true
}