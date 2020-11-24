package main

import "math"

var dp []int
func rob(nums []int) int {
	if len(nums)==0{
		return 0
	}
	dp = make([]int,len(nums))
	for i:=0;i<len(dp);i++{
		dp[i]=-1
	}
	return rob_helper(nums,0)
}
func rob_helper(nums []int,start int) int{
	if start>=len(nums){
		return 0;
	}
	if dp[start]!=-1 {
		return dp[start]
	}
	res:= math.Max(float64(rob_helper(nums,start+1)),float64(nums[start]+rob_helper(nums,start+2)))
	dp[start]=int(res)
	return int(res)
}
func rob_f(nums []int) int{
	n:=len(nums)
	dp_i,dp_i_1,dp_i_2:=0,0,0
	for i:=n-1;i>=0;i--{
		dp_i = int(math.Max(float64(dp_i_1),float64(nums[i]+dp_i_2)))
		dp_i_2=dp_i_1
		dp_i_1=dp_i
	}
	return dp_i
}
