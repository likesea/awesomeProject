package main

import "math"

//一次买卖
//dp[i][1][0] = max(dp[i-1][1][0],dp[i-1][1][1]+prices[i])
// dp[i][1][1] = max(dp[i-1][1][1],dp[i-1][1][0]-prices[i])
//base case
//dp[-1][k][0]=dp[i][0][0]=0
//dp[-1][k][1]=dp[i][0][1]=-infinity
func maxProfit(prices []int) int {
	 n:=len(prices)
	if n==0{
		return 0
	}
	 k:=1
	 dp :=make([][][]int,n)
	 for k0,_:=range dp{
	 	dp[k0]=make([][]int,k+1)
	 	for k1,_:=range dp[k0]{
	 		dp[k0][k1]=make([]int,2)
		}
	 }
	 for i:=0;i<n;i++{
	 	if i-1==-1{
			//dp[-1][k][0]=dp[i][0][0]=0
			//dp[-1][k][1]=dp[i][0][1]=-infinity
			for ko:=0;ko<k+1;ko++{
				dp[0][ko][0]=0
				dp[0][ko][1]=-prices[0]
			}
			for j:=0;j<n;j++{
				dp[j][0][0]=0
				dp[j][0][1]=-prices[j]
			}
		}else{
			for k1:=k;k1>=1;k1--{
				dp[i][k][0] = int(math.Max(float64(dp[i-1][k][0]),float64(dp[i-1][k][1]+prices[i])))
				dp[i][k][1] = int(math.Max(float64(dp[i-1][k][1]),float64(dp[i-1][k-1][0]-prices[i])))
			}
		}

	 }

	return dp[n-1][k][0]

}
