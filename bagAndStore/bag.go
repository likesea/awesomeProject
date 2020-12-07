package bagAndStone

import "math"

func Knapsack(w int, n int, weights []int,values []int) int{
	dp :=init_knap_bag_dp(n+1,w+1)
	for i:=1;i<=n;i++{
		for j:=1;j<=w;j++{
			if  j-weights[i]<0{
				dp[i][j]=dp[i-1][j]
			}else{
				dp[i][j]=int(math.Max(float64(dp[i-1][j-weights[i-1]]+values[i]),float64(dp[i-1][j])))
			}
		}
	}
	return dp[n][w]
}
func init_knap_bag_dp(weight int, bagNum int) [][]int{
	dp:=make([][]int,bagNum)
	for i:=0;i<len(dp);i++{
		dp[i]=make([]int,weight)
	}
	return dp
}

func canPartition(nums []int) bool {
	sum:=0
	for i:=0;i<len(nums);i++{
		sum+=nums[i]
	}
	if sum%2!=0{
		return  false
	}

	sum = sum/2

	dp:=init_canPartition(len(nums)+1,sum+1)
	for i:=0;i<len(nums);i++{
		dp[i][0]=true
	}
	for i:=1;i<=len(nums);i++{
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] < 0 {
				dp[i][j]=dp[i-1][j]
			}else{
				dp[i][j]= dp[i-1][j]||dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][sum]
}
func init_canPartition(N int, W int)[][]bool  {
	m:=make([][]bool,N);
	for i:=0;i<N;i++{
		m[i]=make([]bool,W)
	}
	return m
}
func canPartition2(nums []int) bool {
	sum:=0
	for i:=0;i<len(nums);i++{
		sum+=nums[i]
	}
	if sum%2!=0{
		return  false
	}

	sum = sum/2

	dp:=init_canPartition2(sum+1)
	dp[0]=true
	for i:=0;i<=len(nums);i++{
		//i=0时，即只有一个物品
		//如果此时sum=0，j=0,dp[j]=dp[0]=true
		// 如果此时sum!=0, 此时dp[sum]=false,最终dp[sum]取决于物品重量是否等于sum（因为只有一个物品，即j-nums[0]是否为0）
		for j := sum; j >=0; j-- {
			if j-nums[i] >= 0 {
				dp[j]=dp[j]||dp[j-nums[i]]
			}
		}
	}
	return dp[sum]
}
func init_canPartition2( W int)[]bool  {
	m:=make([]bool,W);
	return m
}
