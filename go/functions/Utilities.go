package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func DateStringToIntSlice(dateString string) ([]int, error) {
	dateParts := strings.Split(dateString, "-")

	var intSlice []int
	for _, part := range dateParts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		intSlice = append(intSlice, num)
	}

	return intSlice, nil
}

func IsNumberinSlice(number int, slice []int) bool {
	for _, value := range slice {
		if value == number {
			return true
		}
	}
	return false
}

func IsStringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func parseDate(dateStr string) string {
	parts := strings.Split(dateStr, "-")
	if len(parts) != 3 {
		return ""
	}
	return parts[2] + "-" + parts[1] + "-" + parts[0]
}

func DateStringToYear(date string) (float64, error) {
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid date format: %s", date)
	}

	year, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse year: %s", parts[2])
	}

	return year, nil
}
