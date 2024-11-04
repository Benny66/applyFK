package holiday

import (
	"log"
	"time"
)

var HolidayMap = map[string][]string{
	"2024": Holidays2024,
}
var WorkWeekMap = map[string][]string{
	"2024": NeedWorkWeek2024,
}

// IsWeekend 判断给定日期是否是周末
func IsWeekend(t time.Time) bool {
	// time.Weekday() 返回一周中的某天，time.Saturday 和 time.Sunday 分别代表周六和周日
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

func IsHoliday(date string) bool {
	currentYear := time.Now().Format("2006") // 获取当前年份
	holidays := HolidayMap[currentYear]
	if holidays == nil {
		log.Fatal("holiday map is nil")
		return false
	}
	for _, holiday := range holidays {
		if holiday == date {
			return true
		}
	}
	t, _ := time.Parse(DayLayout, date)
	if IsWeekend(t) {
		workWeeks := WorkWeekMap[currentYear]
		for _, workWeek := range workWeeks {
			if workWeek == date {
				return false
			}
		}
		return true
	}
	return false
}
