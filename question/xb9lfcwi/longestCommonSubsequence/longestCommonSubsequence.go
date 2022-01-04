package longestCommonSubsequence

import "math"

var dpMemo [][]int

// https://leetcode-cn.com/problems/longest-common-subsequence/
func longestCommonSubsequence(text1 string, text2 string) int {
	dpMemo = make([][]int,len(text1))

	for i:=0;i<len(dpMemo);i++{
		dpMemo[i] = make([]int,len(text2))
	}

	return dp(text1,len(text1)-1,text2,len(text2)-1)
}

func dp(text1 string,tIndex1 int,text2 string ,tIndex2 int) int {

	if tIndex1 <0 || tIndex2<0{
		return 0
	}

	if dpMemo[tIndex1][tIndex2]!=0{
		return  dpMemo[tIndex1][tIndex2]
	}

	if text1[tIndex1] == text2[tIndex2] {
		dpMemo[tIndex1][tIndex2] = dp(text1,tIndex1-1,text2,tIndex2-1)+1
	}else{
		dpMemo[tIndex1][tIndex2] = int(math.Max(float64(dp(text1, tIndex1-1, text2, tIndex2)), float64(dp(text1, tIndex1, text2, tIndex2-1))))
	}

	return dpMemo[tIndex1][tIndex2]

}