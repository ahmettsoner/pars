package time

import (
	"time"
)

func FormatDate(t time.Time, format string) string {
	return t.Format(format)
}

func ParseDate(dateStr string, format string) (time.Time, error) {
	return time.Parse(format, dateStr)
}

func GetCurrentDate() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}

func FormatTime(ti time.Time, format string) string {
	return ti.Format(format)
}

func ParseTime(timeStr string, format string) (time.Time, error) {
	return time.Parse(format, timeStr)
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func AddDuration(d time.Duration) time.Time {
	return time.Now().Add(d)
}
