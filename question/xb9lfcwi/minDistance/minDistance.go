package minDistance

import "math"

func minDistance(word1 string, word2 string) int {
	var mathLen int
	if len(word1) > len(word2) {
		mathLen = len(word1)
	}else{
		mathLen = len(word2)
	}

	dp := make([][]int,mathLen+1)
	for i:=0;i<mathLen+1;i++{
		dp[i] = make([]int,mathLen+1)
	}

	for i := 1;i<len(word1)+1;i++{
		dp[i][0] = i
	}

	for j := 1;j<len(word2)+1;j++{
		dp[0][j] = j
	}

	for i:=0;i<len(word1);i++{
		for j:=0;j<len(word2);j++{
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
				continue
			}

			deleteNum := dp[i+1][j] + 1
			insertNum := dp[i][j+1]+1
			replaceNum := dp[i][j] +1

			dp[i+1][j+1] = int(math.Min(float64(deleteNum), math.Min(float64(insertNum), float64(replaceNum))))
		}
	}

	return dp[len(word1)][len(word2)]
}