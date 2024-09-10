package util

import "time"

const (
	dateTimeJaFormat = "2006年01月02日 03:04 PM"
)

func DateTimeJaFormatter(t time.Time) string {
	return t.Format(dateTimeJaFormat)
}
