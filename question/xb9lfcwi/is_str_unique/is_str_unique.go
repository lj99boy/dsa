package is_str_unique


func IsStrUnique(astr string) bool {
	//println(2&1)
	//println(2064&(1<<4))
	//return false

	var mark uint32 = 0
	for k,_:= range astr {
		offset := astr[k] - 'a'
		if (mark & (1<<offset)) !=0 {
			return false
		} else {
			mark |= 1<<offset
		}

	}

	return true

	//countMap := make(map[string]int)
	//for i:=0;i<len(astr);i++{
	//	if countMap[astr[i:i+1]] >= 1{
	//		return false
	//	}
	//
	//	countMap[astr[i:i+1]]++
	//}
	//
	//return true
}
