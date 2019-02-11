package global_functions

import (
	"strconv"
	"time"
)

func FloatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 2, 64)
}

func IntToString(inputNum int) string {
	return strconv.FormatInt(int64(inputNum), 10)
}

func StringToInt(input string) (int64, error) {
	return strconv.ParseInt(input, 10, 64)
}

func GetCurrentColumn() string {
	column := "C"
	hour := time.Now().Hour()

	if hour >= 0 && hour < 6 {
		column = "C"
	} else if hour >= 6 && hour < 12 {
		column = "D"
	} else if hour >= 12 && hour < 18 {
		column = "E"
	} else if hour >= 18 && hour < 23 {
		column = "F"
	}

	return column
}
