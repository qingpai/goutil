package util

import (
	"github.com/qingpai/goutil/localtime"
	"golang.org/x/exp/constraints"
	"time"
)

func FilterNull(input *int) int {
	if input == nil {
		return 0
	}

	return *input
}

func CalcAge(birthday *localtime.LocalDate) int64 {
	if birthday == nil {
		return 0
	}
	t := birthday.ToTime()
	age := time.Now().Year() - t.Year()
	/*	if now.YearDay() < birthday.YearDay() {
		age--
	}*/
	if age > 150 {
		age = 0
	}
	if age < 1 {
		age = 0
	}

	return int64(age)
}

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
