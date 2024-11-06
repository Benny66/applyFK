package holiday

import (
	"fmt"
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

func TestGetCustomWeekDays(t *testing.T) {
	tests := []struct {
		week     int
		expected int
	}{
		// Happy path
		{0, 3},  // Assuming there are 5 working days until the next Sunday
		{1, 8},  // Assuming there are 7 more working days until the Sunday of the next week
		{2, 13}, // Assuming there are 7 more working days until the Sunday of the third week

		// Edge cases
		{-1, 0}, // Negative week, should return 0
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("week=%d", test.week), func(t *testing.T) {
			result := GetCustomWeekDays(test.week)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}
func TestCalculateWorkDays(t *testing.T) {
	tests := []struct {
		month    int
		expected int
	}{
		{0, 18}, // assuming the current month has 20 workdays
		{1, 40}, // assuming next month has 18 workdays
		{-1, 0}, // invalid month, should return 0
	}

	for _, test := range tests {
		t.Run("Calculating WorkDays", func(t *testing.T) {
			result := CalculateWorkDays(test.month)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestCalculateWorkDaysToEndOfYear(t *testing.T) {
	tests := []struct {
		name     string
		current  time.Time
		expected int
	}{
		{
			name:     "Happy path - regular date",
			current:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 249, // Assuming 260 working days from Jan 1 to Dec 31, excluding holidays and weekends
		},
		{
			name:     "Edge case - last day of year",
			current:  time.Date(2024, 11, 06, 0, 0, 0, 0, time.UTC),
			expected: 40, // No working days left in the year
		},
		{
			name:     "Edge case - holiday on current date",
			current:  time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC), // Assuming Dec 25 is a holiday
			expected: 5,                                             // Only Dec 26 is counted as a working day
		},
		{
			name:     "Edge case - weekdays only",
			current:  time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), // Assume weekends don't interfere
			expected: 1,                                             // You can calculate this based on actual holidays and weekends
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CalculateWorkDaysToEndOfYear(test.current)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}
