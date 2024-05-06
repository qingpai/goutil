package util

import (
	"code.qingpai365.com/erp/goutil/localtime"
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
