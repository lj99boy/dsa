package canPlaceFlowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed)<1{
		return false
	}

	if len(flowerbed)==1{
		if flowerbed[0] == 0 {
			return n <= 1
		}

		if flowerbed[0]==1{
			return n==0
		}
	}

	prev:=0
	target:=0
	next:=0
	canAddPNum :=0

	for i:=0;i<len(flowerbed);i++{
		target = i
		prev=target-1
		next=target+1
		if prev==-1{
			if flowerbed[target]!=1 && flowerbed[next]!=1 {
				flowerbed[target] = 1
				canAddPNum++
			}
		} else if next == len(flowerbed){
			if flowerbed[target]!=1 && flowerbed[prev]!=1 {
				flowerbed[target] = 1
				canAddPNum++
			}
		}else{
			if flowerbed[target]!=1 && flowerbed[next]!=1 && flowerbed[prev]!=1{
				flowerbed[target] = 1
				canAddPNum++
			}
		}
	}

	if canAddPNum>=n{
		return true
	}else{
		return false
	}
}