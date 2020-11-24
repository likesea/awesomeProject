package main

import (
	"math"
	"sort"
)

type intint [][]int

func removeCoveredIntervals(intervals [][]int) int {
	intvs := intint(intervals)
	//先进行排序
	sort.Sort(intvs)
	left,right:= intvs[0][0], intvs[0][1]
	res:=0
	for i:=1;i<len(intvs);i++{
		intv:=intvs[i]
		//覆盖区间
		if left<=intv[0] && right >=intv[1] {
			res++
		}
		//合并区间
		if right>=intv[0]&&right<=intv[1]{
			right = intv[1]
		}
		//不相交
		if right< intv[0]{
			left = intv[0]
			right = intv[1]
		}
	}
	return len(intvs)-res
}

func(s intint) Len()int{
	return len(s)
}
func (s intint) Less(i,j int) bool{
	if(s[i][0]==s[j][0]){
		return s[i][1] >s[j][1]
	}
	return s[i][0] <s[j][0]
}
func (s intint) Swap(i,j int){
	s[i],s[j] =s[j],s[i]
}
func merge(intervals [][]int) [][]int {
	intvs := intint(intervals)
	//先进行排序
	sort.Sort(intvs)
	res :=make([][]int,0)
	res = append(res, intvs[0])
	cur:=res[0]
	for i:=1;i<len(intvs);i++{
		if intvs[i][0] <= cur[1] {
			cur[1] = int(math.Max(float64(intvs[i][1]),float64(cur[1])))
		}
		if intvs[i][0]>cur[1]{
			res = append(res, intvs[i])
			cur = intvs[i]
		}
	}
	return res
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	 i,j:=0,0
	 res:=make([][]int,0)
	 for ; i<len(A) && j <len(B);{
	 	 A1,A2:=A[i][0],A[i][1]
	 	 B1,B2:=B[j][0],B[j][1]
	 	 if A1<=B2 && A2>=B1 {
			 res = append(res, []int{int(math.Max(float64(A1),float64(B1))),int(math.Min(float64(A2),float64(B2)))})
		 }
		 //无交集
		 if B2<A2 {
			j++
		 }else{
		 	i++
		 }
	 }
	 return  res
}