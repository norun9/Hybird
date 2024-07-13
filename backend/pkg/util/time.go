package util

import "time"

const (
	timeOnly12Hr = "2006年01月02日 03:04 PM" // 12時間制のフォーマット
)

func TimeOnly12HrFormatter(t time.Time) string {
	return t.Format(timeOnly12Hr)
}
