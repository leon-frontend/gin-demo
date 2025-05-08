package models

import "time"

// 获取当前日期，年月日格式
func GetDay() string {
	return time.Now().Format("20060102")
}
