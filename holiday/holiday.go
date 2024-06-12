package holiday

import "time"

// IsWeekend 判断给定日期是否是周末
func IsWeekend(t time.Time) bool {
	// time.Weekday() 返回一周中的某天，time.Saturday 和 time.Sunday 分别代表周六和周日
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

func IsHoliday(date string) bool {
	for _, holiday := range Holidays2024 {
		if holiday == date {
			return true
		}
	}
	t, _ := time.Parse(DayLayout, date)
	if IsWeekend(t) {
		for _, holiday := range NeedWorkWeek2024 {
			if holiday == date {
				return false
			}
		}
		return true
	}
	return false
}
