package models

import "time"

// UnixToTime 函数会将时间戳转换为年月日时分秒格式
func UnixToTime(timestamp int) string {
	// 将一个 Unix 时间戳（以秒为单位）转换为 Go 语言中的 time.Time 类型对象。
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
