package main

import (
	"fmt"
	"leet-code-study/Solutions"
)

func main() {
	fmt.Println("hello world")
	Solutions.MinDistance("123", "123")
}

/*
42 接雨水
*/
//func trap(height []int) int {
//	left,right:=0,len(height)-1
//	leftMax,rightMax,res:=0,0,0
//
//	for left<right {
//		if leftMax<height[left]{
//			leftMax=height[left]
//		}
//		if rightMax<height[right]{
//			rightMax=height[right]
//		}
//		if leftMax<rightMax{
//			res+=leftMax-height[left]
//			left++
//		}else{
//			res+=rightMax-height[right]
//			right--
//		}
//
//	}
//	return res
//}
/*
41 缺失的第一个正数
*/

//func firstMissingPositive(nums []int) int {
//	// nums[left] = left + 1
//	left, right := 0, len(nums)
//	for left < right {
//		if nums[left] == left+1 {
//			left += 1
//		} else if nums[left] < left+1 || nums[left] > right || nums[left] == nums[nums[left]-1] {
//			// 出现缺失
//			nums[left], nums[right-1] = nums[right-1], nums[left]
//			right--
//		} else {
//			index := nums[left]
//			nums[left], nums[index-1] = nums[index-1], nums[left]
//		}
//	}
//	return left + 1
//}

/*
40.组合总和 2
*/
//func combinationSum2(candidates []int, target int) [][]int {
//	res:=make([][]int,0)
//	arr:=make([]int,0)
//	tem:=0
//	index:=0
//	ha:=make(map[int]int)
//	for i:=0;i<len(candidates);i++{
//		tem=candidates[i]
//		index=i
//		for j:=i;j<len(candidates);j++ {
//			if tem>candidates[j]{
//				tem=candidates[j]
//				index=j
//			}
//		}
//		candidates[i],candidates[index]=candidates[index],candidates[i]
//	}
//	tem=candidates[0]
//	ha[tem]=1
//	for i:=1;i<len(candidates);i++{
//		if tem==candidates[i]{
//			ha[tem]++
//		}else {
//			tem=candidates[i]
//			ha[tem]++
//		}
//	}
//	res=ana(candidates,target,arr,res,0,ha)
//	return res
//}
//
//func ana(c []int,target int,arr []int ,res [][]int,begin int,ha map[int]int)[][]int{
//	for i:=begin;i<len(c);i+=ha[c[i]]{
//		for j:=0;j<ha[c[i]];j++{
//				target-=(j+1)*c[i]
//				if target<0{
//					target+=(j+1)*c[i]
//					break
//				}
//				for k:=0;k<=j;k++{
//					arr= append(arr, c[i])
//				}
//				if target==0{
//					res= append(res, make([]int,len(arr)))
//					copy(res[len(res)-1], arr)
//					arr=arr[:len(arr)-1-j]
//					target+=(j+1)*c[i]
//					break
//				}else{
//					res=ana(c,target,arr,res,i+ha[c[i]],ha)
//					arr=arr[:len(arr)-1-j]
//					target+=(j+1)*c[i]
//				}
//
//		}
//	}
//	return res
//}

/*
39.组合总和
*/
//func combinationSum(candidates []int, target int) [][]int {
//	res:=make([][]int,0)
//	arr:=make([]int,0)
//	tem:=0
//	index:=0
//	for i:=0;i<len(candidates);i++{
//		tem=candidates[i]
//		index=i
//		for j:=i;j<len(candidates);j++ {
//			if tem>candidates[j]{
//				tem=candidates[j]
//				index=j
//			}
//		}
//		candidates[i],candidates[index]=candidates[index],candidates[i]
//	}
//	//fmt.Println(candidates)
//	res=ana(candidates,target,arr,res,0)
//	return res
//}
//
//func ana(c []int,target int,arr []int ,res [][]int,begin int)[][]int{
//	for i:=begin;i<len(c);i++{
//		if c[i]==target {
//			arr=append(arr, c[i])
//			res = append(res,make([]int,len(arr)))
//			copy(res[len(res)-1], arr)
//			arr=arr[:len(arr)-1]
//		}else{
//			target-=c[i]
//			if target<0{
//				return res
//			}
//			arr= append(arr, c[i])
//			res=ana(c,target,arr,res,i)
//			arr=arr[:len(arr)-1]
//			target+=c[i]
//		}
//	}
//	return res
//}

/*
38,外观数列
*/

//func countAndSay(n int) string {
//	res:=make([]byte,1)
//	res[0]='1'
//	if n==1{
//		return string(res)
//	}
//	var a byte
//	var c byte
//	count:=0
//	tem:=make([]byte,0)
//	for i:=2;i<=n;i++{
//		a =res[0]
//		count=0
//		for _,c=range res {
//			if c==a{
//				count++
//			}else{
//				tem=append(tem, byte(count+48), a)
//				a=c
//				count=1
//			}
//		}
//		tem=append(tem,byte(count+48), a)
//		res=tem
//		tem=make([]byte,0)
//		fmt.Println(string(res))
//	}
//	return string(res)
//}

/*
37.解数独
*/

//func solveSudoku(board [][]byte)  {
//	min:=[]int{9,9,9,9,9,9,9,9,9}
//	for i:=0;i<9;i++{
//		for j:=0;j<9;j++{
//			if board[i][j]!='.'{
//				min[i]--
//			}
//		}
//	}
//	fmt.Println(fill(board,min))
//	fmt.Println(board)
//}
//
//func fill(b [][]byte,min []int)bool  {
//	index:=0
//	tem:=10
//	for i,va:=range min{
//		if va!=0&&va <tem{
//			tem=va
//			index=i
//		}
//	}
//	min[index]--
//	if min[index]==0{
//		fmt.Println(min)
//	}
//	if tem==10{
//		return true
//	}
//	cl:=0
//	for i:=0;i<9;i++{
//		if b[index][i]=='.'{
//			cl=i
//			break
//		}
//	}
//	var i byte
//	ofC:=(cl/3)*3
//	ofR:=(index/3)*3
//
//	for i=49;i<58;i++{
//		j:=0
//		for ;j<9;j++{
//			if b[index][j]==i||b[j][cl]==i{
//				break
//			}
//			if j<3{
//				if b[ofR][ofC+j]==i{
//					break
//				}
//
//			}else if j<6{
//				if b[ofR+1][ofC+j-3]==i{
//					break
//				}
//			}else{
//				if b[ofR+2][ofC+j-6]==i{
//					break
//				}
//			}
//		}
//		if j!=9{
//			continue
//		}
//		b[index][cl]=i
//		if fill(b,min){
//			return true
//		}else{
//			b[index][cl]='.'
//		}
//	}
//	min[index]++
//	return false
//}

/*
36.有效的数独
*/
//func isValidSudoku(board [][]byte) bool {
//	r:=make([][]byte,9)
//	l:=make([][]byte,9)
//	b:=make([][]byte,9)
//	bIndex:=0
//	var tem byte
//	for i:=0;i<9;i++{
//		for j:=0;j<9;j++{
//			if board[i][j]=='.'{
//				continue
//			}
//			for  _,tem=range r[i]{
//				if tem==board[i][j]{
//					return false
//				}
//			}
//			r[i]= append(r[i], board[i][j])
//			for  _,tem=range l[j]{
//				if tem==board[i][j]{
//					return false
//				}
//			}
//			l[j]= append(l[j], board[i][j])
//			bIndex=(i/3)*3+j/3
//			for  _,tem=range b[bIndex]{
//				if tem==board[i][j]{
//					return false
//				}
//			}
//			b[bIndex]= append(b[bIndex], board[i][j])
//		}
//	}
//	return true
//}

/*
35
*/
//func searchInsert(nums []int, target int) int {
//	l,r,mid:=0,len(nums),0
//	for l<r{
//		mid=l+(r-l)/2
//		if nums[mid]>=target{
//			r=mid
//		}else{
//			l=mid+1
//		}
//	}
//	return l
//}

/*
34
*/
//func searchRange(nums []int, target int) []int {
//	res := []int{-1, -1}
//	if len(nums) == 0 {
//		return res
//	}
//	if len(nums) == 1 {
//		if nums[0] == target {
//			res[0] = 0
//			res[1] = 0
//		}
//		return res
//	}
//	l, r, mid := 0, len(nums)-1, 0
//	for l < r {
//		mid = l + (r - l)/2
//		if nums[mid] >= target {
//			r = mid
//		} else {
//			l = mid + 1
//		}
//	}
//	if nums[l] == target {
//		res[0] = l
//	} else {
//		return res
//	}
//	r=len(nums)-1
//	for l < r {
//		mid=l+(r-l)/2
//		if nums[mid]<=target{
//			l=mid+1
//		}else{
//			r=mid
//		}
//	}
//	res[1]=l
//	return res
//
//}

/*
33 搜索旋转排序数组 二分查找
*/
//func search(nums []int, target int) int {
//	if len(nums) == 0 {
//		return -1
//	} else if len(nums) == 1 {
//		if nums[0] != target {
//			return -1
//		} else {
//			return 0
//		}
//	}
//	l, r, mid := 0, len(nums)-1, 0
//	for {
//		if r-l<=2{
//			for i:=l;i<=r;i++{
//				if nums[i]==target {
//					return i
//				}
//			}
//			return -1
//		}
//		if nums[l] < nums[r] {
//			mid = l + (r-l)/2
//			if nums[mid] == target {
//				return mid
//			} else if nums[mid] > target {
//				r = mid
//			} else {
//				l = mid
//			}
//		} else {
//			mid = l + (r-l)/2
//			if nums[l]<nums[mid]{
//				if nums[l]<=target&&nums[mid]>=target{
//					r=mid
//				}else{
//					l=mid
//				}
//			}else{
//				if nums[mid]<=target&&nums[r]>=target{
//					l=mid
//				}else {
//					r=mid
//				}
//			}
//
//		}
//	}
//
//}

/*
32 最长有效括号 动态规划
*/
//func longestValidParentheses(s string) int {
//	if len(s)==0||len(s)==1{
//		return 0
//	}
//	dp:=make([]int,len(s))
//	dp[0]=0
//	max:=0
//	if s[0]=='('&&s[1]==')'{
//		dp[1]=2
//		max=2
//	}else{
//		dp[1]=0
//	}
//
//	for i:=2;i<len(s);{
//		if s[i]==')'{
//			if s[i-1]=='('{
//				dp[i]=dp[i-2]+2
//				if dp[i]>max{
//					max=dp[i]
//				}
//				i++
//
//			}else{
//				if i-1-dp[i-1]>=0&&s[i-1-dp[i-1]]=='('{
//					if i-2-dp[i-1]>=0{
//						dp[i]=dp[i-1]+2+dp[i-2-dp[i-1]]
//					}else{
//						dp[i]=dp[i-1]+2
//					}
//
//					if dp[i]>max{
//						max=dp[i]
//					}
//				}
//				i++
//			}
//		}else{
//			i++
//		}
//	}
//	return max
//
//
//}

/*
32 栈
*/
//func longestValidParentheses(s string) int {
//	l:=len(s)
//	if l==0||l==1{
//		return 0
//	}
//	max:=0
//	stack:=make([]byte,0)
//	a:=make([]int,0)
//	b:=make([]int8,l)
//	sIndex:=0
//	sMax:=0
//	for i:=0;i<l;i++{
//		if s[i]=='('{
//			if sIndex>=sMax{
//				stack=append(stack,s[i])
//				a=append(a,i)
//				sMax++
//			}else{
//				stack[sIndex]=s[i]
//				a[sIndex]=i
//			}
//			sIndex++
//		}else{
//			if sIndex==0{
//			}else{
//				if stack[sIndex-1]=='('{
//					b[a[sIndex-1]],b[i]=1,1
//					sIndex--
//				}
//			}
//		}
//	}
//	tem:=0
//	b= append(b, 0)
//	for i:=0;i<l+1;{
//		if b[i]==1{
//			i+=2
//			tem+=2
//		}else{
//			if tem>max{
//				max=tem
//			}
//			tem=0
//			i++
//		}
//	}
//	return max
//}

/*

28 实现strStr（）
*/
//func strStr(haystack string, needle string) int {
//	l:=len(needle)
//	if l==0{
//		return 0
//	}
//	next:=make([]int,l,l)
//	next[0]=-1
//	j:=0
//	if l==2{
//		next[1]=0
//	}else{
//		for i:=1;i<l-1;{
//			if j==-1||needle[i]==needle[j]{
//				j++
//				next[i+1]=j
//				i++
//			}else{
//				j=next[j]
//			}
//		}
//	}
//	l2:=len(haystack)
//	j=0
//	res:=-1
//	for i:=0;i<l2;{
//		if j==-1||haystack[i]==needle[j]{
//			if j==l-1{
//				res=i-l+1
//				break
//			}
//			j++
//			i++
//		}else{
//			j=next[j]
//		}
//	}
//	return res
//}

/*
29 两数相除
*/
//func divide(dividend int, divisor int) int {
//	var flag =1
//	if dividend<=0{
//		dividend=-dividend
//		flag=-flag
//	}
//	if divisor<=0{
//		divisor=-divisor
//		flag=-flag
//	}
//	if dividend>math.MaxInt32{
//		if divisor==1&&flag==-1{
//			return -math.MaxInt32-1
//		}else if divisor==1&&flag==1{
//			return math.MaxInt32
//		}
//
//	}
//	if dividend==math.MaxInt32&&divisor==1{
//		if flag==1{
//			return math.MaxInt32
//		}else{
//			return -math.MaxInt32
//		}
//	}
//	var times,ri,res,l,r,mid int
//	res=0
//	l=0
//	mid=0
//	for dividend>=divisor {
//		times=0
//		if l==0{
//			ri=divisor
//			for ;ri<=dividend;times++{
//				ri<<=1
//			}
//			l=times
//			r=0
//			dividend-=ri>>1
//		}else{
//			for{
//				mid=(l-r)>>1+r
//				ri=divisor<<mid
//				if ri==dividend{
//					times=mid+1
//					dividend=0
//					break
//				}else if ri>dividend{
//					l=mid
//					continue
//				}else{
//					if (ri<<1)>dividend{
//						dividend-=ri
//						times=mid+1
//						l=times-1
//						r=0
//						break
//					}else{
//						r=mid+1
//					}
//				}
//			}
//		}
//		res +=1<<(times-1)
//	}
//	if flag==1{
//		return res
//	}else{
//		return -res
//	}
//
//
//}

/*
31 下一个排列
*/

//func nextPermutation(nums []int)  {
//	l:= len(nums)
//	i:=l-1
//	for ;i-1>=0;i--{
//		if nums[i]>nums[i-1] {
//			j:=i
//			for ;j<l;j++{
//				if nums[j]<=nums[i-1]{
//					break
//				}
//			}
//			nums[i-1],nums[j-1]=nums[j-1],nums[i-1]
//
//			n:=i
//			m:=l-1
//			for ;m>n;{
//				nums[n],nums[m]=nums[m],nums[n]
//				n++
//				m--
//			}
//			return
//		}
//	}
//	n:=0
//	m:=l-1
//	for ;m>n;{
//		nums[n],nums[m]=nums[m],nums[n]
//		n++
//		m--
//	}
//	return
//
//
//
//}
/*
30 串联子串
*/

//func findSubstring(s string, words []string) []int {
//	l:=len(words)
//	res:=make([]int,0)
//	if l==0{
//		return res
//	}
//	var ok bool
//	aMap:=make(map[string]int)
//	for i:=0;i<l;i++{
//		_,ok:=aMap[words[i]]
//		if ok{
//			aMap[words[i]]++
//		}else{
//			aMap[words[i]]=1
//		}
//
//	}
//	wL:= len(words[0])
//	index:=0
//	sMap:=make(map[string]int)
//	for ;index<=len(s)-l*wL;index++{
//		for i:=0;i<l;i++{
//			sMap[words[i]]=0
//		}
//		for i:=0;i<l*wL;i+=wL{
//			_,ok=sMap[s[index+i:index+wL+i]]
//			if !ok{
//				break
//			}else{
//				sMap[s[index+i:index+wL+i]]++
//				if i==(l-1)*wL{
//					sw:=0
//					for key,value:=range sMap{
//						if value!=aMap[key]{
//							sw=1
//						}
//					}
//					if sw==0{
//						res= append(res, index)
//					}
//				}
//
//			}
//		}
//
//	}
//	return res
//}
//
