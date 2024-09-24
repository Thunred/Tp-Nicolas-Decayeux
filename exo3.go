package exam

import "sort"

func Ft_non_overlap(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	if len(intervals) == 0 {
		return count
	}

	endLastInterval := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		currentInterval := intervals[i]

		if currentInterval[0] < endLastInterval {
			count++
		} else {
			endLastInterval = currentInterval[1]
		}
	}

	return count
}
