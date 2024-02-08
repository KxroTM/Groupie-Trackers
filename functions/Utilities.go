package Groupie_Trackers

import (
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
