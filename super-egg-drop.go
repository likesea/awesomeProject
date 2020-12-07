package main

import "math"

type int2 struct {
	K int
	N int
}
var m map[int2]int

func superEggDrop(K int, N int) int {
	m = make(map[int2]int)
	r:= sed(K,N)
	return r
}

func sed (K int, N int) int {
	if K==1{
		return N
	}
	if N==0 {
		return 0
	}
	if val,ok:=m[int2{K,N}];ok{
		return val
	}
	var resEgg =1000000
	//for i:=1;i<N+1;i++{
	//	resEgg = int(math.Min(float64(resEgg),math.Max(float64(sed(K-1,i-1)),float64(sed(K,N-i)))+1))
	//}
	lo,hi := 1,N
	for ; lo<=hi;{
		mid := (lo+hi)/2
		broken :=sed(K-1,mid-1)
		not_broken := sed(K,N-mid)
		if broken > not_broken{
			hi = mid - 1
			resEgg = int(math.Min(float64(resEgg), float64(broken + 1)))
		}else{
			lo = mid + 1
			resEgg = int(math.Min(float64(resEgg), float64(not_broken  + 1)))
		}
	}
	m[int2{K,N}]=resEgg
	return resEgg
}
