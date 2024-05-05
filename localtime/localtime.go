package localtime

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

const DateFormat = "2006-01-02"
const TimeFormat = "2006-01-02 15:04:05"

type LocalDate time.Time
type LocalTime time.Time

/// local date

func (t *LocalDate) UnmarshalJSON(data []byte) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")

	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalDate(time.Time{})
		return
	}

	// 指定解析的格式
	now, err := time.ParseInLocation(`"`+DateFormat+`"`, string(data), loc)
	*t = LocalDate(now)
	return
}

func (t LocalDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(DateFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, DateFormat)
	b = append(b, '"')
	return b, nil
}

// Value 写入 mysql 时调用
func (t LocalDate) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Local().Format(DateFormat)), nil
}

// Scan 检出 mysql 时调用
func (t *LocalDate) Scan(v interface{}) error {
	var tTime time.Time
	tTime, _ = dateparse.ParseAny(fmt.Sprintf("%s", v))

	*t = LocalDate(tTime)
	return nil
}

// 用于 fmt.Println 和后续验证场景
func (t LocalDate) String() string {
	return time.Time(t).Format(DateFormat)
}

/// local time

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")

	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	// 指定解析的格式
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), loc)
	*t = LocalTime(now)
	return
}

func (t *LocalDate) IsZero() bool {
	return time.Time(*t).IsZero()
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) Date() LocalDate {
	return LocalDate(time.Time(t))
}

// Value 写入 mysql 时调用
func (t LocalTime) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Local().Format(TimeFormat)), nil
}

// Scan 检出 mysql 时调用
func (t *LocalTime) Scan(v interface{}) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	//tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	var tTime time.Time
	tTime, _ = dateparse.ParseAny(fmt.Sprintf("%s", v))

	*t = LocalTime(tTime)
	return nil
}

// 用于 fmt.Println 和后续验证场景
func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

func NowDate() *LocalDate {
	return MakeLocalDate(time.Now())
}

func Now() *LocalTime {
	return MakeLocalTime(time.Now())
}

func MakeLocalDate(input time.Time) *LocalDate {
	localDate := LocalDate(input)

	return &localDate
}

func MakeLocalTime(input time.Time) *LocalTime {
	localDate := LocalTime(input)

	return &localDate
}

func (t LocalDate) ToTime() time.Time {
	return time.Time(t)
}

func (t LocalTime) ToTime() time.Time {
	return time.Time(t)
}
