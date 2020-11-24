package main


func missingNumber(nums []int) int{
	left:=0
	right :=len(nums)
	var mid int
	for ;left<right ;  {
		mid=left+(right-left)/2
		if nums[mid]==mid{
			left = mid+1
		}else if nums[mid]>mid {
			right = mid
		}else if nums[mid]<mid{
			//不可能的情况
			continue
		}
	}
	return right
}
//func missingNumber(nums []int) int {
//	length := len(nums)
//	var sum int
//	for i:=0;i<=length;i=i+1{
//		sum+=i
//	}
//	for _,n:=range nums{
//		sum-=n
//	}
//	return sum
//}
