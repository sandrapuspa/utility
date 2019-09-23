package utility

func SumCalculation(values ...int64) int64 {
	result := int64(0)

	for _, v := range values {
		result += v
	}

	return result
}
