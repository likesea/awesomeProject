package main

import "fmt"

var res [][]int = make([][]int,0)
func main() {
	fmt.Println(threeSumTarget([]int{-1, 0, 1, 2, -1, -4},0))
}

func permute(nums []int,track []int){
	if len(track)==len(nums){
		res = append(res, track)
		return
	}
	for i:=0;i<len(nums);i++{
		if contains(track,nums[i]){
			continue
		}
		track = append(track,nums[i])
		permute(nums,track)
		track = track[:len(track)-1]
	}
}
func contains(s []int,num int) bool{
	for _,item:=range s{
		if item==num{
			return true
		}
	}
	return false
}