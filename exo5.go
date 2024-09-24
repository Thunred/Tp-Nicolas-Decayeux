package exam

func Ft_max_substring(s string) int {
	Pos := make(map[rune]int)
	max := 0
	debut := 0
	for i, char := range s {
		if precedent, found := Pos[char]; found {

			if precedent >= debut {
				debut = precedent + 1
			}
		}
		Pos[char] = i
		chaincara := i - debut + 1
		if chaincara > max {
			max = chaincara
		}
	}
	return max
}
