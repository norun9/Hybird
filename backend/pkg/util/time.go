package util

import "time"

const (
	timeOnly12Hr = "3:04 PM" // 12時間制のフォーマット
)

func TimeOnly12HrFormatter(t time.Time) string {
	return t.Format(timeOnly12Hr)
}
