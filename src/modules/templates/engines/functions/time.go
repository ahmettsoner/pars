package functions

import (
	"time"

	_time "parsdevkit.net/core/utils/time"
)

type TimeFuncs struct{}

func (d TimeFuncs) FormatDate(t time.Time, format string) string {
	return _time.FormatDate(t, format)
}

func (d TimeFuncs) ParseDate(dateStr string, format string) (time.Time, error) {
	return _time.ParseDate(dateStr, format)
}

func (d TimeFuncs) GetCurrentDate() time.Time {
	return _time.GetCurrentDate()
}

func (t TimeFuncs) FormatTime(ti time.Time, format string) string {
	return _time.FormatTime(ti, format)
}

func (t TimeFuncs) ParseTime(timeStr string, format string) (time.Time, error) {
	return _time.ParseTime(timeStr, format)
}

func (t TimeFuncs) GetCurrentTime() time.Time {
	return _time.GetCurrentTime()
}

// func (t TimeFuncs) AddDuration(d time.Duration) time.Time {
// 	return time.Now().Add(d)
// }
