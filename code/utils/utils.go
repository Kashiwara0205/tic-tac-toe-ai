package utils

func GetMaxIdx(slice []int) int {
	max := -9999999
	maxIdx := -1
	for idx, s := range slice {
		if max < s {
			max = s
			maxIdx = idx
		}
	}

	return maxIdx
}

func GetMinIdx(slice []int) int {
	min := 9999999
	minIdx := -1
	for idx, s := range slice {
		if min > s {
			min = s
			minIdx = idx
		}
	}

	return minIdx
}
