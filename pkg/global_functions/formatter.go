package global_functions

import (
	"strconv"
)

func FloatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 2, 64)
}

func IntToString(inputNum int) string {
	return strconv.FormatInt(int64(inputNum), 10)
}

func StringToFloat(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}
