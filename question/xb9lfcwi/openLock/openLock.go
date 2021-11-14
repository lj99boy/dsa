package openLock

func getNowCanReachLock(nowStr string) (res []string)  {
	for i:=0;i<len(nowStr);i++{
		tStr := nowStr[i]
		var tStrUp byte
		if tStr == '9' {
			tStrUp = '0'
		}else{
			tStrUp = tStr+1
		}

		var tStrDown byte
		if tStr == '0' {
			tStrDown = '9'
		}else{
			tStrDown = tStr-1
		}

		tResUp := nowStr[:i] + string(tStrUp) + nowStr[i+1:]
		tResDown := nowStr[:i] + string(tStrDown) + nowStr[i+1:]

		res = append(res,tResUp)
		res = append(res,tResDown)
	}

	return res
}

func isInSlice(items []string,target string)bool {
	for i:=0;i<len(items);i++{
		if target == items[i]{
			return true
		}
	}
	return false
}

func openLock(deadends []string, target string) int {

	if target == "" {
		return -1
	}

	bFSHolder := make([]string,0)

	bFSHolder = append(bFSHolder,"0000")

	visited := make(map[string]struct{})

	visited["0000"] = struct{}{}

	step :=0
	for ;len(bFSHolder)>0; {
		tLen := len(bFSHolder)
		for i:=0;i<tLen;i++{
			item := bFSHolder[i]
			if item == target{
				return step
			}

			if isInSlice(deadends,item) {
				continue
			}

			rawNextItems := getNowCanReachLock(item)
			nextItems := make([]string,0)

			for j:=0;j<len(rawNextItems);j++{
				if _,ok:=visited[rawNextItems[j]];!ok{
					visited[rawNextItems[j]] = struct{}{}
					nextItems = append(nextItems,rawNextItems[j])
				}
			}

			bFSHolder = append(bFSHolder,nextItems...)
		}
		bFSHolder = bFSHolder[tLen:]
		step++
	}

	return -1
}
