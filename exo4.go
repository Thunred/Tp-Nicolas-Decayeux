package exam

func Ft_profit(prices []int) int {
	var result int
	if len(prices) <= 1 {
		return 0
	}
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			benef := prices[i] - min
			if benef > result {
				result = benef
			}
		}
	}

	return result
}
