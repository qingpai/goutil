package util

func FilterNull(input *int) int {
	if input == nil {
		return 0
	}

	return *input
}
