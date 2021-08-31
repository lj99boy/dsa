package solve_n_queen

//func SolveNQueens(n int) [][]string {
//	var res [][]string
//
//	var closure func([]string)
//
//	closure = func(tClassBoard []string) {
//		if len(tClassBoard) == n+1 {
//			res = append(res, tClassBoard[:len(tClassBoard)-1])
//			return
//		}
//
//		for ; ; {
//			var tmpBoard = tClassBoard
//			tmpBoard = getNextRow(tmpBoard)
//			if tmpBoard != nil {
//				closure(append(tmpBoard, strings.Repeat(".", n)))
//				tmpBoard = tmpBoard[:len(tmpBoard)-1]
//			}
//		}
//	}
//
//	return res
//}
//
//func getNextRow(input []string) []string {
//	handleRow := input[len(input)-1]
//
//	for k,c := range input[len(input)-1]{
//		if c == '.'{
//
//		} else {
//
//		}
//	}
//}
