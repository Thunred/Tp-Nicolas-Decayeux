package exam

func Ft_missing(num []int) int {
	if min(num) != 0 && min(num) > 0 {
		return 0
	}
	max := max(num)
	n := len(num)
	if max != n {
		return -1
	}
	expectedSum := (n * (n + 1)) / 2 // somme des eniter de 0 a n
	actualSum := 0

	for _, v := range num {
		actualSum += v
	}

	return expectedSum - actualSum
}

func max(arr []int) int {
	if len(arr) == 0 {
		panic("Le tableau est vide")
	}

	maxVal := arr[0] // on suppose que le premier élément est le max
	for _, value := range arr {
		if value > maxVal {
			maxVal = value
		}
	}
	return maxVal
}

func min(arr []int) int {
	if len(arr) == 0 {
		panic("Le tableau est vide")
	}

	maxVal := arr[0] // on suppose que le premier élément est le max
	for _, value := range arr {
		if value < maxVal {
			maxVal = value
		}
	}
	return maxVal
}
