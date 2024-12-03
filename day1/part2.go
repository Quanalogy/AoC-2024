package main

func TotalSimilarity(a, b []int) int {
	res := 0

	for i := 0; i < len(a); i++ {
		value := a[i]
		count := 0
		for j := 0; j < len(b); j++ {
			if value == b[j] {
				count++
			}
		}
		res += value * count
	}

	return res
}
