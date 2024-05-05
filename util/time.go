package util

import "time"

func GetTimePeriod(time time.Time) string {
	switch time.Hour() {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		return "上午"
	case 12, 13, 14, 15, 16, 17:
		return "下午"
	case 18, 19, 20, 21, 22, 23:
		return "晚上"
	}

	return ""
}

func TimestampFormat(timestamp int64) string {
	if timestamp < 1 {
		return ""
	}
	date := time.Unix(timestamp, 0)
	return date.Format("2006-01-02 15:04:05")
}

func Clock(date time.Time, hour int, minute int, second int) time.Time {
	year, month, day := date.Date()
	return time.Date(year, month, day, hour, minute, second, 0, date.Location())
}

func HourEqual(date1 time.Time, date2 time.Time) bool {
	return date1.Year() == date2.Year() && date1.Month() == date2.Month() && date1.Day() == date2.Day() && date1.Hour() == date2.Hour()
}

func DayEqual(date1 time.Time, date2 time.Time) bool {
	return date1.Year() == date2.Year() && date1.Month() == date2.Month() && date1.Day() == date2.Day()
}

func HourEqualByTimestamp(date1 int64, date2 time.Time) bool {
	return HourEqual(time.Unix(date1, 0).UTC().Local(), date2)
}

func DayEqualByTimestamp(date1 int64, date2 time.Time) bool {
	return DayEqual(time.Unix(date1, 0).UTC().Local(), date2)
}
