package holiday

import "time"

// 定义原始时间格式
var DayLayout string = "2006/01/02"
var TimeLayout string = "2006/01/02 15:04"

func CurrentDayStr(day int) string {
	// 获取当前时间
	currentTime := time.Now().AddDate(0, 0, day)
	// 格式化为 YYYY-MM-DD 格式
	return currentTime.Format("2006/01/02")
}

func CurrentDayTime() time.Time {
	// 获取当前时间
	currentTime := time.Now()
	// 格式化为 YYYY-MM-DD 格式
	currentDate := currentTime.Format("2006-01-02")
	// 解析时间字符串
	parsedTime, _ := time.Parse("2006-01-02", currentDate)
	return parsedTime
}

func ParsedTime(date string) (time.Time, error) {
	return time.Parse(TimeLayout, date)
}
