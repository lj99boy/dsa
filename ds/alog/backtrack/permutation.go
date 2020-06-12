package backtrack

type SelectedIntArr []int

func (arr SelectedIntArr) Contains(i int) bool {
	for _, val := range arr {
		if val == i {
			return true
		}
	}
	return false
}

var res []SelectedIntArr

func MainGetPermutation(in []int) []SelectedIntArr {
	selected := SelectedIntArr{}

	getPermutation(in, selected)

	return res
}

func getPermutation(nums []int, selected SelectedIntArr) {
	if len(selected) == len(nums) {
		res = append(res, selected)
		return
	}

	for i := 0; i < len(nums); i++ {
		if selected.Contains(nums[i]) {
			continue
		}
		selected = append(selected, nums[i])
		getPermutation(nums, selected)

		selected = selected[:len(selected)-1]
	}

}
