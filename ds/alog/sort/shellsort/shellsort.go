//希尔排序
//是插入排序的改进，在直接进行插入排序之前，先进行几轮不同跨度（不同希尔增量）的排序（其实直接插入排序也就是希尔增量为1的特殊情况）
//https://baijiahao.baidu.com/s?id=1644158198885715432&wfr=spider&for=pc
package shellsort

func ShellSort(arr []int) []int {
	d := len(arr)
	//希尔增量分区,这里的d>0,d=d/2可以根据增量序列改变
	for ; d > 0; {
		d = d / 2
		//根据希尔增量划分排序组
		for i := 0; i < d; i++ {
			//直接插入排序
			for j := i; j <= len(arr)-1; j += d {
				for innerJ := j + d; innerJ > j && innerJ <= (len(arr)-1); innerJ -= d {
					if arr[innerJ] < arr[innerJ-d] {
						arr[innerJ], arr[innerJ-d] = arr[innerJ-d], arr[innerJ]
					}
				}
			}
		}
	}
	return arr
}
