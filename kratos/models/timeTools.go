package models

import (
	"time"
)

func TimeToUnix(t string) int64 { // 时间----->时间戳
	const layout = "2006-01-02 15:04:05"
	Unix, _ := time.Parse(layout, t)
	return Unix.Unix()
}

func UnixToTime(timeUnix int) string { // 时间戳----->时间
	t := time.Unix(int64(timeUnix), 0)
	return t.Format(time.RFC3339)
}
