package global_functions

import (
	"strconv"
)

func FloatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 2, 64)
}
