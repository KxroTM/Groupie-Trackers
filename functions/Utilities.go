package functions

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
