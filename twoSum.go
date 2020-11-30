package main

import "sort"

func twoSumTarget(nums []int, start int, target int) [][]int {
  length:=len(nums)
  sort.Ints (nums)
  i,j:=start,length-1
  res:=make([][]int,0)
  for;i<j;{
  	sum :=nums[i]+nums[j]
  	left,right:=nums[i],nums[j]
  	if sum<target{
  		for ;i<j &&nums[i]==left; { i++}
	} else if sum>target{
		for ;i<j &&nums[j]==right; { j--}
	}else {
		res = append(res,[]int{nums[i],nums[j]})
		for ;i<j &&nums[i]==left; { i++}
		for ;i<j &&nums[j]==right; { j--}
	  }
  }
	return res
}
func indexOf(element int, data []int) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
}

func threeSumTarget(nums []int,target int) [][]int {
	sort.Ints(nums)
	length:=len(nums)
	res := make([][]int,0)
	for i:=0;i<length;i++{
		r := twoSumTarget(nums,i+1,target-nums[i])
		if r!=nil && len(r)!=0{
			for j:=0;j<len(r);j++{
				r[j] = append(r[j],nums[i])
			}
			res= append(res, r...)
		}
		for ;i<length-1 &&nums[i]==nums[i+1]; {
			i++
		}
	}
	return res
}