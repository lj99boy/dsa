package tmp1

import (
	"fmt"
	"sort"
)

var resHolder []string

func demo()  {
	fmt.Println("a"[0])
	fmt.Println("A"[0])
	fmt.Println("z"[0])
	fmt.Println("Z"[0])

}

func FindSnakeStr()  {
	rawStr:=""
	fmt.Scan(&rawStr)

	resHolder = make([]string,0)
	Exec(rawStr)
}

func Exec(str string)  {
	//字符元素栈
	minStack:=make([]int,0)
	//元素关联下标
	tMap:=make(map[uint8][]int)

	for i:=0;i<len(str);i++{
		if str[i]>=65 && str[i]<=122{
			minStack = append(minStack,int(str[i]))
			tMap[str[i]] = append(tMap[str[i]],i)
		}
	}

	sort.Ints(minStack)

	for i:=0;i<len(minStack);i++{
		if len(tMap[uint8(minStack[i])])>1{
			if _,ok:=tMap[uint8(minStack[i]+32)];ok && len(tMap[uint8(minStack[i]+32)])>1{

			}
		}
	}
}
