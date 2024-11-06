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

// 判断是否是周日
func IsSunday(date string) bool {
	t, _ := time.Parse(DayLayout, date)
	return t.Weekday() == time.Sunday
}

// 距离指定周日还有多天工作日，week=1为当前周日，week=2为下周日，以此类推
func GetCustomWeekDays(week int) int {
	if week < 0 {
		return 0
	}
	hadWeek := 0
	start := 0
	days := 0
	for {
		currentDate := CurrentDayStr(start)
		if !IsHoliday(currentDate) {
			days++
		}
		isSunday := IsSunday(currentDate)
		if isSunday {
			hadWeek++
			if hadWeek > week {
				break
			}
		}
		start++
	}
	return days
}

// calculateWorkDays 计算从今天起，到指定月份的最后一天的工作日天数
func CalculateWorkDays(month int) int {
	if month < 0 {
		return 0
	}
	current := time.Now()
	// 时间为当天的0点
	current = time.Date(current.Year(), current.Month(), current.Day(), 0, 0, 0, 0, time.Local)
	targetMonth := current.AddDate(0, month, 0)
	lastDayOfMonth := getLastDayOfMonth(targetMonth.Year(), int(targetMonth.Month()))

	days := 0
	for t := current; t.Before(lastDayOfMonth) || t.Equal(lastDayOfMonth); t = t.Add(24 * time.Hour) {
		tStr := t.Format(DayLayout)
		if !IsHoliday(tStr) {
			days++
		}
	}
	return days
}

// getLastDayOfMonth 获取指定年份和月份的最后一天日期
func getLastDayOfMonth(year, month int) time.Time {
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
}

// CalculateWorkDaysToEndOfYear 计算从今天起，到今年的最后一天的工作日天数
func CalculateWorkDaysToEndOfYear(current time.Time) int {
	// 时间为当天的0点
	current = time.Date(current.Year(), current.Month(), current.Day(), 0, 0, 0, 0, time.Local)
	endOfYear := time.Date(current.Year(), 12, 31, 0, 0, 0, 0, time.UTC)

	days := 0
	for t := current; t.Before(endOfYear) || t.Equal(endOfYear); t = t.Add(24 * time.Hour) {
		tStr := t.Format(DayLayout)
		if !IsHoliday(tStr) {
			days++
		}
	}
	return days
}
