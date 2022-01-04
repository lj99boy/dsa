package hitBallon

func hitBallon(ballons [][]int) int {
	for i:=len(ballons)-1;i>0;i--{
		for j:=0;j<i;j++{
			if ballons[j][1]>ballons[j+1][1]{
				ballons[j],ballons[j+1]=ballons[j+1],ballons[j]
			}
		}
	}

	sum :=1
	tBallon := ballons[0]
	for i:=1;i<len(ballons);i++{
		if !isInRange(tBallon[1],ballons[i]){
			tBallon = ballons[i]
			sum++
		}
	}
	return sum
}

func isInRange(val int,ballon []int) bool {
	if val <= ballon[1] && val>=ballon[0]{
		return true
	}else{
		return false
	}

}