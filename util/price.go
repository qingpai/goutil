package util

import "github.com/shopspring/decimal"

func ParsePrice(input string) int {
	result, err := decimal.NewFromString(input)
	if err != nil {
		return 0
	}

	return int(result.Round(2).Mul(decimal.NewFromInt32(100)).IntPart())
}
