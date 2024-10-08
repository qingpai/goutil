package datetime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// A Date represents a date (year, month, day).
//
// This type does not include location information, and therefore does not
// describe a unique 24-hour timespan.
type Date struct {
	Year  int        // Year (e.g., 2014).
	Month time.Month // Month of the year (January = 1, ...).
	Day   int        // Day of the month, starting at 1.
}

// DateOf returns the Date in which a time occurs in that time's location.
func DateOf(t time.Time) Date {
	var d Date
	d.Year, d.Month, d.Day = t.Date()
	return d
}

// ParseDate parses a string in RFC3339 full-date format and returns the date value it represents.
func ParseDate(s string) (Date, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return Date{}, err
	}
	return DateOf(t), nil
}

// String returns the date in RFC3339 full-date format.
func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

// IsValid reports whether the date is valid.
func (d Date) IsValid() bool {
	return DateOf(d.In(time.UTC)) == d
}

// In returns the time corresponding to time 00:00:00 of the date in the location.
//
// In is always consistent with time.Date, even when time.Date returns a time
// on a different day. For example, if loc is America/Indiana/Vincennes, then both
//
//	time.Date(1955, time.May, 1, 0, 0, 0, 0, loc)
//
// and
//
//	civil.Date{Year: 1955, Month: time.May, Day: 1}.In(loc)
//
// return 23:00:00 on April 30, 1955.
//
// In panics if loc is nil.
func (d Date) In(loc *time.Location) time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, loc)
}

// AddDays returns the date that is n days in the future.
// n can also be negative to go into the past.
func (d Date) AddDays(n int) Date {
	return DateOf(d.In(time.UTC).AddDate(0, 0, n))
}

// DaysSince returns the signed number of days between the date and s, not including the end day.
// This is the inverse operation to AddDays.
func (d Date) DaysSince(s Date) (days int) {
	// We convert to Unix time so we do not have to worry about leap seconds:
	// Unix time increases by exactly 86400 seconds per day.
	deltaUnix := d.In(time.UTC).Unix() - s.In(time.UTC).Unix()
	return int(deltaUnix / 86400)
}

// Before reports whether d occurs before d2.
func (d Date) Before(d2 Date) bool {
	if d.Year != d2.Year {
		return d.Year < d2.Year
	}
	if d.Month != d2.Month {
		return d.Month < d2.Month
	}
	return d.Day < d2.Day
}

// After reports whether d occurs after d2.
func (d Date) After(d2 Date) bool {
	return d2.Before(d)
}

// IsZero reports whether date fields are set to their default value.
func (d Date) IsZero() bool {
	return (d.Year == 0) && (int(d.Month) == 0) && (d.Day == 0)
}

// MarshalText implements the encoding.TextMarshaler interface.
// The output is the result of d.String().
func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The date is expected to be a string in a format accepted by ParseDate.
func (d *Date) UnmarshalText(data []byte) error {
	var err error
	*d, err = ParseDate(string(data))
	return err
}

// Value 写入 mysql 时调用
func (t Date) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01" {
		return nil, nil
	}
	return []byte(t.String()), nil
}

// Scan 检出 mysql 时调用
func (t *Date) Scan(v interface{}) error {
	var err error
	input := ""
	switch v.(type) {
	case []uint8:
		vv := v.([]uint8)
		if vv == nil || len(vv) == 0 {
			t = nil
			return nil
		}
		input = b2s(vv)
	case string:
		input = v.(string)
	}
	if *t, err = ParseDate(input); err != nil {
		return err
	}
	return nil
}

// A Time represents a time with nanosecond precision.
//
// This type does not include location information, and therefore does not
// describe a unique moment in time.
//
// This type exists to represent the TIME type in storage-based APIs like BigQuery.
// Most operations on Times are unlikely to be meaningful. Prefer the DateTime type.
type Time struct {
	Hour   int // The hour of the day in 24-hour format; range [0-23]
	Minute int // The minute of the hour; range [0-59]
}

// TimeOf returns the Time representing the time of day in which a time occurs
// in that time's location. It ignores the date.
func TimeOf(t time.Time) Time {
	var tm Time
	tm.Hour, tm.Minute, _ = t.Clock()
	return tm
}

// ParseTime parses a string and returns the time value it represents.
// ParseTime accepts an extended form of the RFC3339 partial-time format. After
// the HH:MM:SS part of the string, an optional fractional part may appear,
// consisting of a decimal point followed by one to nine decimal digits.
// (RFC3339 admits only one digit after the decimal point).
func ParseTime(s string) (Time, error) {
	t, err := time.Parse("15:04", s)
	if err != nil {
		return Time{}, err
	}
	return TimeOf(t), nil
}

// String returns the date in the format described in ParseTime. If Nanoseconds
// is zero, no fractional part will be generated. Otherwise, the result will
// end with a fractional part consisting of a decimal point and nine digits.
func (t Time) String() string {
	return fmt.Sprintf("%02d:%02d", t.Hour, t.Minute)
}

// IsValid reports whether the time is valid.
func (t Time) IsValid() bool {
	// Construct a non-zero time.
	tm := time.Date(2, 2, 2, t.Hour, t.Minute, 0, 0, time.UTC)
	return TimeOf(tm) == t
}

// IsZero reports whether time fields are set to their default value.
func (t Time) IsZero() bool {
	return (t.Hour == 0) && (t.Minute == 0)
}

// Before reports whether t occurs before t2.
func (t Time) Before(t2 Time) bool {
	if t.Hour != t2.Hour {
		return t.Hour < t2.Hour
	}
	if t.Minute != t2.Minute {
		return t.Minute < t2.Minute
	}

	return false
}

// After reports whether t occurs after t2.
func (t Time) After(t2 Time) bool {
	return t2.Before(t)
}

func (t Time) AfterOrEqual(t2 Time) bool {
	if t.Hour == t2.Hour && t.Minute == t2.Minute {
		return true
	}
	return t2.Before(t)
}

func (t Time) AddHour(n int) Time {
	return Time{Hour: t.Hour + n, Minute: t.Minute}
}

func (t Time) AddMinute(n int) Time {
	tm := time.Date(2, 2, 2, t.Hour, t.Minute, 0, 0, time.UTC)
	return TimeOf(tm.Add(time.Duration(n) * time.Minute))
}

func (t Time) CountHalfHour(t2 Time) int {
	if t2.Before(t) {
		return 0
	}

	count := 0
	for start := t; start.Before(t2); start = start.AddMinute(30) {
		count++
	}

	return count
}

// MarshalText implements the encoding.TextMarshaler interface.
// The output is the result of t.String().
func (t Time) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time is expected to be a string in a format accepted by ParseTime.
func (t *Time) UnmarshalText(data []byte) error {
	var err error
	*t, err = ParseTime(string(data))
	return err
}

// Value 写入 mysql 时调用
func (t Time) Value() (driver.Value, error) {
	return []byte(t.String()), nil
}

// Scan 检出 mysql 时调用
func (t *Time) Scan(v interface{}) error {
	var err error

	input := ""
	switch v.(type) {
	case []uint8:
		vv := v.([]uint8)
		if vv == nil || len(vv) == 0 {
			t = nil
			return nil
		}
		input = b2s(vv)
	case string:
		input = v.(string)
	}

	if *t, err = ParseTime(input); err != nil {
		return err
	}
	return nil
}

// A DateTime represents a date and time.
//
// This type does not include location information, and therefore does not
// describe a unique moment in time.
type DateTime struct {
	Date Date
	Time Time
}

// Note: We deliberately do not embed Date into DateTime, to avoid promoting AddDays and Sub.

// DateTimeOf returns the DateTime in which a time occurs in that time's location.
func DateTimeOf(t time.Time) DateTime {
	return DateTime{
		Date: DateOf(t),
		Time: TimeOf(t),
	}
}

// ParseDateTime parses a string and returns the DateTime it represents.
// ParseDateTime accepts a variant of the RFC3339 date-time format that omits
// the time offset but includes an optional fractional time, as described in
// ParseTime. Informally, the accepted format is
//
//	YYYY-MM-DDTHH:MM:SS[.FFFFFFFFF]
//
// where the 'T' may be a lower-case 't'.
func ParseDateTime(s string) (DateTime, error) {
	t, err := time.Parse("2006-01-02T15:04", s)
	if err != nil {
		t, err = time.Parse("2006-01-02t15:04", s)
		if err != nil {
			return DateTime{}, err
		}
	}
	return DateTimeOf(t), nil
}

// String returns the date in the format described in ParseDate.
func (dt DateTime) String() string {
	return dt.Date.String() + "T" + dt.Time.String()
}

// IsValid reports whether the datetime is valid.
func (dt DateTime) IsValid() bool {
	return dt.Date.IsValid() && dt.Time.IsValid()
}

// In returns the time corresponding to the DateTime in the given location.
//
// If the time is missing or ambigous at the location, In returns the same
// result as time.Date. For example, if loc is America/Indiana/Vincennes, then
// both
//
//	time.Date(1955, time.May, 1, 0, 30, 0, 0, loc)
//
// and
//
//	civil.DateTime{
//	    civil.Date{Year: 1955, Month: time.May, Day: 1}},
//	    civil.Time{Minute: 30}}.In(loc)
//
// return 23:30:00 on April 30, 1955.
//
// In panics if loc is nil.
func (dt DateTime) In(loc *time.Location) time.Time {
	return time.Date(dt.Date.Year, dt.Date.Month, dt.Date.Day, dt.Time.Hour, dt.Time.Minute, 0, 0, loc)
}

// Before reports whether dt occurs before dt2.
func (dt DateTime) Before(dt2 DateTime) bool {
	return dt.In(time.UTC).Before(dt2.In(time.UTC))
}

// After reports whether dt occurs after dt2.
func (dt DateTime) After(dt2 DateTime) bool {
	return dt2.Before(dt)
}

// IsZero reports whether datetime fields are set to their default value.
func (dt DateTime) IsZero() bool {
	return dt.Date.IsZero() && dt.Time.IsZero()
}

// MarshalText implements the encoding.TextMarshaler interface.
// The output is the result of dt.String().
func (dt DateTime) MarshalText() ([]byte, error) {
	return []byte(dt.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The datetime is expected to be a string in a format accepted by ParseDateTime
func (dt *DateTime) UnmarshalText(data []byte) error {
	var err error
	*dt, err = ParseDateTime(string(data))
	return err
}

// Value 写入 mysql 时调用
func (t DateTime) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(t.String()), nil
}

// Scan 检出 mysql 时调用
func (t *DateTime) Scan(v interface{}) error {
	var err error
	input := ""
	switch v.(type) {
	case []uint8:
		vv := v.([]uint8)
		if vv == nil || len(vv) == 0 {
			t = nil
			return nil
		}
		input = b2s(vv)
	case string:
		input = v.(string)
	}
	if *t, err = ParseDateTime(input); err != nil {
		return err
	}
	return nil
}

func b2s(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}
