package generateParenthesis

import (
	"fmt"
	"strings"
)

//https://leetcode-cn.com/problems/generate-parentheses/

var ansStrArr []string

func generateParenthesis(n int) []string {

	ansStrArr = make([]string,0)

	doRec(make([]string,0),0,0,n)

	fmt.Println(ansStrArr)

	return ansStrArr
}

func doRec(nowStrArr []string,left,right,n int)  {
	if right > left {
		return
	}

	if left>n{
		return
	}

	if right == left && n == right{
		ansStrArr = append(ansStrArr,strings.Join(nowStrArr,""))
	}

	doRec(append(nowStrArr,"("),left+1,right,n)

	doRec(append(nowStrArr,")"),left,right+1,n)
}