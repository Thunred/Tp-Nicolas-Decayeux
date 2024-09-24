package exam

type Interval struct {
	Debut int
	Fin   int
}

func Ft_non_overlap(intervals []Interval) int {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i].Fin > intervals[j].Fin {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
	count := 0
	FindernierInterval := intervals[0].Fin
	for i := 1; i < len(intervals); i++ {
		ActuInterval := intervals[i]
		if ActuInterval.Debut < FindernierInterval {
			count++
		} else {
			FindernierInterval = ActuInterval.Fin
		}
	}

	return count
}
