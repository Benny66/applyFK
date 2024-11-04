package holiday

import (
	"testing"
	"time"
)

// Mock holidays and work weeks for testing

func TestIsHoliday(t *testing.T) {
	// Set the current year for the test
	tests := []struct {
		date     string
		expected bool
	}{
		// Happy path cases
		{"2024/01/01", true},  // New Year's Day
		{"2024/12/25", false}, // Christmas Day
		{"2024/01/02", false}, // Not a holiday, but a work week

		// Weekend test cases
		{"2024/01/06", true},  // Saturday
		{"2024/01/07", true},  // Sunday
		{"2024/01/08", false}, // Monday (not a holiday)

		// Edge cases
		{"2024/12/31", false},   // Not a holiday (New Year's Eve)
		{"invalid-date", false}, // Invalid date format
	}

	for _, tt := range tests {
		t.Run(tt.date, func(t *testing.T) {
			got := IsHoliday(tt.date)
			if got != tt.expected {
				t.Errorf("IsHoliday(%v) = %v; want %v", tt.date, got, tt.expected)
			}
		})
	}
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		date     time.Time
		expected bool
	}{
		{time.Date(2023, 10, 14, 0, 0, 0, 0, time.UTC), true},  // Saturday
		{time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC), true},  // Sunday
		{time.Date(2023, 10, 16, 0, 0, 0, 0, time.UTC), false}, // Monday
		{time.Date(2023, 10, 17, 0, 0, 0, 0, time.UTC), false}, // Tuesday
		{time.Date(2023, 10, 18, 0, 0, 0, 0, time.UTC), false}, // Wednesday
		{time.Date(2023, 10, 19, 0, 0, 0, 0, time.UTC), false}, // Thursday
		{time.Date(2023, 10, 20, 0, 0, 0, 0, time.UTC), false}, // Friday
	}

	for _, test := range tests {
		result := IsWeekend(test.date)
		if result != test.expected {
			t.Errorf("IsWeekend(%v) = %v; expected %v", test.date, result, test.expected)
		}
	}
}

func TestCurrentDayStr(t *testing.T) {
	tests := []struct {
		day      int
		expected string
	}{
		{0, time.Now().Format("2006/01/02")},                        // Happy path: current day
		{1, time.Now().AddDate(0, 0, 1).Format("2006/01/02")},       // Next day
		{-1, time.Now().AddDate(0, 0, -1).Format("2006/01/02")},     // Previous day
		{365, time.Now().AddDate(0, 0, 365).Format("2006/01/02")},   // One year later
		{-365, time.Now().AddDate(0, 0, -365).Format("2006/01/02")}, // One year ago
	}

	for _, test := range tests {
		result := CurrentDayStr(test.day)
		if result != test.expected {
			t.Errorf("CurrentDayStr(%d) = %s; expected %s", test.day, result, test.expected)
		}
	}
}
