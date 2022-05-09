package homework

func average(input [15]float32, N int64) (result float32) {

	result = 0.0
	for _, v := range input {
		result += v
	}

	return result / float32(N)
}
