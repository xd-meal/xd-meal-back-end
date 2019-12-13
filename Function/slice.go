package Function

func MaxString(slice []string) string {
	max := slice[0]
	for _, v := range slice {
		if max < v {
			max = v
		}
	}
	return max
}

func MinString(slice []string) string {
	min := slice[0]
	for _, v := range slice {
		if min > v {
			min = v
		}
	}
	return min
}
