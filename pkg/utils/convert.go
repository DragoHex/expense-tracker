package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ConvertTimetToString converts int month & year to
// named_month-year
// example m=1 & y=2000 --> January-2000
func ConvertTimetToString(m, y int) string {
	return fmt.Sprintf("%s-%d", time.Month(m).String(), y)
}

// SplitStringToTime does the following conversion
// January-2000 --> 1 2000
func SplitStringToTime(monthYear string) (int, int, error) {
	monthYearSlice := strings.Split(monthYear, "-")

	layout := "January"
	t, err := time.Parse(layout, monthYearSlice[0])
	if err != nil {
		return 0, 0, err
	}

	y, err := strconv.Atoi(monthYearSlice[1])
	if err != nil {
		return 0, 0, err
	}

	return int(t.Month()), y, nil
}
