package helper

import (
	"time"
)

func FormatDate(date string) string {

	t, _ := time.Parse("2006-01-02", date)
	return t.Format("02.01.2006")
}
