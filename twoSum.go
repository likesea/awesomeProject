package main

import "sort"

func twoSum(nums []int, target int) []int {
  length:=len(nums)
  intvs:=make([]int,len(nums))
  copy(intvs,nums)
  sort.Ints (nums)
  i,j:=0,length-1
  res:=make([]int,0)
  for;i<j;{
  	if nums[i]+nums[j]==target{
		res =append(res, nums[i],nums[j])
		break
	}
	if nums[i]+nums[j]>target{
		j--
	}
	  if nums[i]+nums[j]<target{
		  i++
	  }
  }
  if len(res)!=0{
  	return []int{indexOf(res[0],intvs),indexOf(res[1],intvs)}
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