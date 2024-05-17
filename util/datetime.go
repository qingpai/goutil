package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"time"

	"github.com/jinzhu/now"
)

func DateEqualUnix(input time.Time, target int64) bool {
	return FormatDate(input) == FormatDate(time.Unix(target, 0))
}

func DateEqual(input time.Time, target time.Time) bool {
	return FormatDate(input) == FormatDate(target)
}

func FormatDate(input time.Time) string {
	return input.Format("2006-01-02")
}

func FormatDateTime(input time.Time) string {
	return input.Format("2006-01-02 15:04:05")
}

// ParseDateTimeAsTimestamp 解析日期时间为时间戳
func ParseDateTimeAsTimestamp(input string) (int64, error) {
	t, err := now.Parse(input)
	if err != nil {
		return 0, err
	}

	return t.Unix(), nil
}

// ParseDateTime 解析日期时间
func ParseDateTime(input string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")

	t, err := now.ParseInLocation(loc, input)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
func ParseDateTimeWithUTC(input string) (time.Time, error) {
	loc, _ := time.LoadLocation("UTC")

	t, err := now.ParseInLocation(loc, input)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
func ParseDateTimeWithFormat(input string, format string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")

	config := &now.Config{
		TimeLocation: loc,
		TimeFormats:  []string{format},
	}

	t, err := config.Parse(input)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func ParseDate(input string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	input = input + " 00:00:00"

	return now.ParseInLocation(loc, input)
}

// ParseStartEndDate 解析开始日期和结束日期
func ParseStartEndDate(c *gin.Context) (time.Time, time.Time, error) {
	startDateParam := c.DefaultQuery("startDate", "")
	endDateParam := c.DefaultQuery("endDate", "")

	if startDateParam == "" || endDateParam == "" {
		return time.Time{}, time.Time{}, errors.New("没有开始日期或结束日期")
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")

	startDate, err := now.ParseInLocation(loc, startDateParam)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("开始日期格式错误")
	}

	endDate, err := now.ParseInLocation(loc, endDateParam)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("结束日期格式错误")
	}

	if startDate.After(endDate) {
		return time.Time{}, time.Time{}, errors.New("开始日期不能大于结束日期")
	}

	return startDate, endDate, nil
}

// ParseStartEndTime 解析开始时间和结束时间
func ParseStartEndTime(c *gin.Context) (time.Time, time.Time, error) {
	startTimeParam := c.DefaultQuery("startTime", "")
	endTimeParam := c.DefaultQuery("endTime", "")

	if startTimeParam == "" || endTimeParam == "" {
		return time.Time{}, time.Time{}, errors.New("没有开始时间或结束时间")
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")

	startDate, err := now.ParseInLocation(loc, startTimeParam)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("开始时间格式错误")
	}

	endDate, err := now.ParseInLocation(loc, endTimeParam)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("结束时间格式错误")
	}
	if DayEqual(endDate, time.Now()) {
		endDate = time.Now()
	} else {
		endDate = endDate.AddDate(0, 0, 1)
	}

	if startDate.After(endDate) {
		return time.Time{}, time.Time{}, errors.New("开始时间不能大于结束时间")
	}

	return startDate, endDate, nil
}

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
